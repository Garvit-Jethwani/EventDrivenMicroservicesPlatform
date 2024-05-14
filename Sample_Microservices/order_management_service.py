
import json
import pika

def publish_event():
    connection = pika.BlockingConnection(pika.ConnectionParameters('localhost'))
    channel = connection.channel()
    channel.queue_declare(queue='orders')
    order = {'id': 123, 'item': 'Test product', 'quantity': 1}
    channel.basic_publish(exchange='', routing_key='orders', body=json.dumps(order))
    print("Order event published")
    connection.close()

publish_event()
