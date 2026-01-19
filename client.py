import requests
import random
import time
import json
import sys

# Configuration
SERVER_URL = "http://localhost:8080"
VECTOR_DIM = 1536 
TOTAL_VECTORS = 50000 

def generate_random_vector(dim):
    """Generates a random list of floats to simulate an embedding."""
    return [random.random() for _ in range(dim)]

def insert_data():
    print(f"[INFO] Starting generation and insertion of {TOTAL_VECTORS} vectors...")
    
    start_time = time.time()
    
    for i in range(TOTAL_VECTORS):
        doc_id = f"doc_{i}"
        vector = generate_random_vector(VECTOR_DIM)
        
        payload = {
            "id": doc_id,
            "vector": vector
        }
        
        try:
            response = requests.post(f"{SERVER_URL}/insert", json=payload)
            response.raise_for_status()
            
            # Print progress every 1000 items to keep stdout clean
            if i % 1000 == 0:
                print(f"    - Processed {i} vectors...")
                
        except Exception as e:
            print(f"[ERROR] Failed to insert {doc_id}: {e}")
            return

    duration = time.time() - start_time
    print(f"[SUCCESS] Insertion complete. Total time: {duration:.2f}s")

def search_data():
    print("\n[INFO] Executing search query...")
    
    
    query_vector = generate_random_vector(VECTOR_DIM)
    
    payload = {
        "vector": query_vector,
        "k": 5 
    }
    
    start_time = time.time()
    
    try:
        response = requests.post(f"{SERVER_URL}/search", json=payload)
        response.raise_for_status()
        results = response.json()
        
    except Exception as e:
        print(f"[ERROR] Search request failed: {e}")
        return

    total_latency = (time.time() - start_time) * 1000 # Convert to ms
    
    print(f"[SUCCESS] Search completed.")
    print(f"    - Client Latency:   {total_latency:.2f}ms")
    print(f"    - Server Processing: {results['latency']}")
    
    print("\nTop 5 Results:")
    print(json.dumps(results['results'], indent=2))

if __name__ == "__main__":
    # Check server health before starting
    try:
        # We expect a 404 because GET / is not defined, but it proves the port is open
        requests.get(SERVER_URL) 
    except requests.exceptions.ConnectionError:
        print("[ERROR] Could not connect to Hermes server at :8080.")
        print("        Please ensure the Go server is running.")
        sys.exit(1)

    insert_data()
    search_data()