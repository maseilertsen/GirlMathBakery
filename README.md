# GirlMathBakery
## Background
This is what happens when you give in to your absolutely fantastic girlfrien

The core concept is that we've bought a kitchen machine, and use girl-math,
> "Just think about how much money we'll save!" - enthusiastic Girlfriend.
---                                                                         

# Usage

## Running Locally
1. Install dependencies:
   ```bash
   go mod download
   ```

2. Set the required environment variable:
   ```bash
   export BAKERY_TOKEN='supersecret'
   ```

3. Run the application:
   ```bash
   go run main.go
   ```

The server will start on `http://localhost:8080`

## Deployment
For simplicity the application is hosted on a local machine in an active she

The `POST` request is sent to the server (local machine) via an `Apple Short
                                                                            
---                                                                         
## Prime database
POST-request to `/seed` with wanted items + cost for baking (unit_cost) vs b
```json                                                                     
[                                                                           
    {"item":"boller","unit_cost":5,"unit_store":60,"unit":""},              
    {"item":"brød","unit_cost":20,"unit_store":40,"unit":""},               
    {"item":"naan","unit_cost":10,"unit_store":20,"unit":""},               
    {"item":"kake","unit_cost":70,"unit_store":175,"unit":""},              
    {"item":"muffins","unit_cost":5,"unit_store":50,"unit":""}              
]                                                                           
```                                                                         
# TODO
- [ ] Deploy via docker
- [ ] Safeguard POST in '/bake' to notify when a new item has been added.   