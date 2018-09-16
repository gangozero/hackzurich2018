class Api {
    constructor() {
        let url = "ws://localhost:8080/ws";
        //let url = "ws://10.15.5.4:8080/ws";
        this.socket = new WebSocket(url, "ws");
        this.socket.onmessage = function (event) {
            try {
                let result = JSON.parse(event.data);
                window.platform.informCentralUsers(result);
            } catch (e) {
                console.log("invalid incomming data" + e);
                console.log(event.data);
            }
        }
    }

    sendPeople(platform, people) {

        const platformWidth = platform.size.width;

        for (const i in people) {
            const p = people[i];
            const data = {
                action: "location",
                payload: {
                    user_id: p.id,
                    major: 100,
                    minor: this.getPlatformZone(platformWidth, p.x)
                }
            };
            let stringData = JSON.stringify(data);
            console.log(stringData);
            this.socket.send(stringData);
        }

        this.socket.send(JSON.stringify(
            {"action": "calibrate", "payload": {"scene": 1}}
        ));
    }

    getPlatformZone(platformWidth, xLocation) {
        let part = platformWidth / 3;
        if (xLocation < part) {
            return 33;
        } else if (xLocation > part * 2) {
            return 31;
        } else {
            return 32;
        }
    }
}