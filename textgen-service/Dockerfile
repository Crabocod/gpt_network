FROM python:3.9-slim

WORKDIR /app
COPY requirements.txt requirements.txt
RUN pip install --no-cache-dir -r requirements.txt

COPY . .
ENV PYTHONPATH=/app/proto
CMD ["python", "main.py"]

EXPOSE 50051