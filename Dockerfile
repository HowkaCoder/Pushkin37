# Use the official KrakenD image
FROM devopsfaith/krakend:latest

# Copy the KrakenD configuration
COPY krakend.json /etc/krakend/krakend.json

# Set the working directory
WORKDIR /app

# Run KrakenD with the provided configuration
CMD ["krakend", "run", "-c", "/etc/krakend/krakend.json"]
