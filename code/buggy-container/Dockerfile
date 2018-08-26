# GNU Compiler Collection (GCC) image
FROM gcc:5

# Copy source to compile in container
COPY . /usr/src/badfibonacci

# Set working directory
WORKDIR /usr/src/badfibonacci

# Compile CPP code
RUN g++ -o badfibonacci badfibonacci.cpp

# Run memory-leaking container
ENTRYPOINT ["./badfibonacci"]