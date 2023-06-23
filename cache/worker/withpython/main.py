import redis

# Connect to Redis
client = redis.Redis(host='localhost', port=6379, decode_responses=True)

# Set a key-value pair
client.set('mykey', 'myvalue')

# Get the value for a key
value = client.get('mykey')
print(value)