import grpc
from concurrent import futures
import service_pb2
import service_pb2_grpc
from transformers import AutoTokenizer, AutoModelForCausalLM
import torch

# Массив для сопоставления названия модели с её путём
MODEL_PATHS = {
    "ЕваGPT": "./model",
}

# Кэш для загруженных моделей
loaded_models = {}

def load_model(model_name):
    """Загружает модель и токенизатор по названию модели."""
    if model_name in loaded_models:
        return loaded_models[model_name]

    if model_name not in MODEL_PATHS:
        raise ValueError(f"Модель с названием '{model_name}' не найдена.")

    model_path = MODEL_PATHS[model_name]
    try:
        tokenizer = AutoTokenizer.from_pretrained(model_path)
        model = AutoModelForCausalLM.from_pretrained(model_path)
        loaded_models[model_name] = (tokenizer, model)
        print(f"Модель '{model_name}' загружена из {model_path}.")
        return tokenizer, model
    except Exception as e:
        raise RuntimeError(f"Ошибка при загрузке модели '{model_name}': {e}")

class TextGenService(service_pb2_grpc.TextGenServiceServicer):
    def GenerateText(self, request, context):
        question = request.question
        model_name = request.model_name

        try:
            tokenizer, model = load_model(model_name)
        except ValueError as ve:
            context.set_code(grpc.StatusCode.NOT_FOUND)
            context.set_details(str(ve))
            return service_pb2.GenerateResponse(answer="")
        except RuntimeError as re:
            context.set_code(grpc.StatusCode.INTERNAL)
            context.set_details(str(re))
            return service_pb2.GenerateResponse(answer="")

        input_text = f"{question} [SEP]"

        # Генерация ответа
        tokens = tokenizer.encode(input_text, return_tensors="pt")
        with torch.no_grad():
            output = model.generate(
                tokens,
                repetition_penalty=1.2,
                do_sample=True,
                top_k=30,
                top_p=0.8,
                temperature=0.8,
                no_repeat_ngram_size=2,
                max_new_tokens=22,
                eos_token_id=tokenizer.eos_token_id,
            )
        response_text = tokenizer.decode(output[0], skip_special_tokens=True)
        answer = response_text.split("[SEP]")[-1].strip()
        return service_pb2.GenerateResponse(answer=answer)

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    service_pb2_grpc.add_TextGenServiceServicer_to_server(TextGenService(), server)
    server.add_insecure_port('[::]:50051')
    print("Сервер Python запущен на порту 50051...")
    server.start()
    server.wait_for_termination()

if __name__ == "__main__":
    serve()