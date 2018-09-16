class Histogram {
    draw(people) {
        let list = {};
        for (let p of people) {

            if(!p.visible) continue;

            if (!list.hasOwnProperty(p.x)) {
                list[p.x] = 1;
            } else {
                list[p.x]++;
            }
        }

        for (var property in list) {
            if (list.hasOwnProperty(property)) {
                let height = list[property];
                let color;

                if (height >= 7) {
                    color = "#ff0000";
                } else if (height >= 5 && height < 7) {
                    color = "#ff8b00";
                } else {
                    color = "#0cff61";
                }


                this.drawRect(parseInt(property) + 1, 70, 1, height * -1 * 2, color);
            }
        }
    }

    drawRect(x, y, width, height, color) {
        x *= gridSize;
        y *= gridSize;
        width *= gridSize;
        height *= gridSize;

        ctx.fillStyle = color;
        ctx.fillRect(x, y, width, height);
    }
}