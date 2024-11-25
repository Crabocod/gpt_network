import grpc
from concurrent import futures
import service_pb2
import service_pb2_grpc
from transformers import AutoTokenizer, AutoModelForCausalLM

# Загрузка токенизатора и модели
tokenizer = AutoTokenizer.from_pretrained("./model")
model = AutoModelForCausalLM.from_pretrained("./model")

class TextGenService(service_pb2_grpc.TextGenServiceServicer):
    def GenerateText(self, request, context):
        question = request.question
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
                early_stopping=True,
                eos_token_id=tokenizer.eos_token_id,
            )
        response_text = tokenizer.decode(output[0], skip_special_tokens=True)
        answer = response_text.split("[SEP]")[-1].strip()
        return service_pb2.GenerateResponse(answer=answer)

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    service_pb2_grpc.add_TextGenServiceServicer_to_server(TextGenService(), server)
    server.add_insecure_port('[::]:50051')
    server.start()
    server.wait_for_termination()

if __name__ == "__main__":
    serve()