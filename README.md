# chatroom-server [Not implemented yet]

# Tech Stack
- Golang
  - github.com/gorilla/websocket
  - github.com/dancannon/gorethink
- Rethink DB


## Testing web socket handling
```javascript
let msg = {
 name: 'channel add',
  data: {
    name: 'hardware support'
  }
}

let subMsg = {
  name: 'channel subscribe'
}

let ws = new WebSocket('ws://localhost:9090');
ws.onopen = () => {
  ws.send(JSON.stringify(subMsg))
  ws.send(JSON.stringify(msg))
}

ws.onmessage = (e) => {
  console.log(e.data);
}
```