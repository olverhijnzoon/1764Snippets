FROM python:3.9-slim
WORKDIR /app
COPY . /app
EXPOSE 14269
ENV NAME headlessService
CMD ["python", "-m", "http.server", "14269"]