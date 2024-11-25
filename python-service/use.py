import torch
from transformers import AutoTokenizer, AutoModelForCausalLM

tokenizer = AutoTokenizer.from_pretrained("./model")
model = AutoModelForCausalLM.from_pretrained("./model")

def generate_response(question):
    test = tokenizer.encode(question, return_tensors="pt")

    with torch.no_grad():
        output = model.generate(
            test,
            repetition_penalty=6.0,
            do_sample=True,
            top_k=5,
            top_p=0.95,
            temperature=1,
            no_repeat_ngram_size=2
        )

    response = tokenizer.decode(output[0], skip_special_tokens=True)
    return response.split("Ответ: ")[-1].strip()
while True:
    question = input("Input: ")
    if question.lower() == 'exit':
        break
    answer = generate_response(question)
    print(f"Answer: {answer}")