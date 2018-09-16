class Person {

    constructor(x, y) {
        this.initX = x;
        this.initY = y;
        this.x = x;
        this.y = y;
        this.completeRandom = true;
        this.locked = false;
        this.visible = false;
        this.custom = false;
        this.boarded = false;
        this.image = null;
        this.hasPopup = false;
        this.trainStatic = false;
        this.id = "user-" + (++window.idCounter);
        this.distributionLocation = 0;
        this.distributionFinished = false;
        this.wentToFirstClass = false;
    }

    loadImage(path) {
        let self = this;
        var img = new Image();
        img.src = path;
        img.onload = function () {
            self.image = img;
        }
    }

    walkAround(distribution, platform) {

        if (distribution && this.custom) {

            if (this.x !== this.distributionLocation) {
                if (this.x < this.distributionLocation) {
                    this.x++;
                    //this.x += Math.random() > 0.5 ? 1 : 0;
                } else if (this.x > this.distributionLocation) {
                    this.x--;
                    //this.x += Math.random() > 0.5 ? -1 : 0;
                }

                if (this.x === this.distributionLocation) {
                    this.distributionFinished = true;
                }
            }

        } else {
            if (this.x !== this.initX || this.y !== this.initY) {
                this.x = this.initX;
                this.y = this.initY;
            } else {
                // this.walkRandom(platform)
            }
        }
    }

    walkRandom(platform) {
        const newX = this.getRandomLoc(this.x, 1, platform.size.width - 1);
        const newY = this.getRandomLoc(this.y, 1, platform.size.height - 1);
        this.x = newX;
        this.y = newY;
    }

    getRandomLoc(current, min, max) {
        const newVal = Math.round(Math.random()) * (Math.random() > 0.5 ? 1 : -1) + current;
        if (newVal >= max) return current;
        else if (newVal < min) return current;
        else return newVal;
    }
}