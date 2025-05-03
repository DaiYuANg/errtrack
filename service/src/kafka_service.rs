use std::time::Duration;
use rdkafka::{producer::FutureProducer, ClientConfig};
use rdkafka::consumer::{BaseConsumer, Consumer};
use rdkafka::error::KafkaResult;
use sea_orm::DatabaseConnection;

// 异步生产者
pub fn create_producer(brokers: &str) -> FutureProducer {
    ClientConfig::new()
        .set("bootstrap.servers", brokers)
        .set("message.timeout.ms", "5000")
        .create()
        .expect("Producer creation error")
}

// 消费者组
pub async fn run_consumer(brokers: &str, topic: &str, db: DatabaseConnection) {
    let consumer: BaseConsumer = ClientConfig::new()
        .set("bootstrap.servers", brokers)
        .set("group.id", "event-processor")
        .set("enable.partition.eof", "false")
        .create()
        .expect("Consumer creation failed");

    consumer.subscribe(&[topic]).expect("Subscribe failed");

    loop {
        match consumer.poll(Duration::from_secs(1)) {
            // Ok(Some(msg)) => {
            //     if let Some(payload) = msg.payload() {
            //         if let Ok(event) = serde_json::from_slice::<Event>(payload) {
            //             if let Err(e) = process_event(&db, event).await {
            //                 eprintln!("Error processing event: {}", e);
            //             }
            //         }
            //     }
            //     consumer.commit_message(&msg, CommitMode::Async).unwrap();
            // }
            // Err(e) => eprintln!("Kafka error: {}", e),
            // _ => {}
            None => {}
            Some(_) => {}
        }
    }
}