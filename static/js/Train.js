class Train {

    constructor(
        inX,
        stopX,
        outX,
        y,
        width,
        height
    ) {
        this.inX = inX;
        this.stopX = stopX;
        this.outX = outX;
        this.x = inX;
        this.y = y;
        this.width = width;
        this.height = height;
        this.people = [];
        this.staticPeople = [];
    }

    arrive() {


        this.addArrivals();

        const self = this;
        const intId = setInterval(function () {
            self.x += 1;//
            if (self.x === self.stopX) {
                clearInterval(intId);
                window.state = window.StateTrainStopped
            }
        }, window.refreshRate);
    }

    depart() {

        const self = this;
        const i = setInterval(function () {

            self.x += 1;
            for (const p in self.people) {
                const person = self.people[p];
                person.x += 1;
            }

            if (self.x === self.outX) {
                self.x = self.inX;
                clearInterval(i);
                self.people = [];
                window.state = window.StateTrainLeft
            }
        }, window.refreshRate);
    }

    board() {
        const self = this;
        const intId = setInterval(function () {

            var moved = false;

            for (var i in self.people) {
                const p = self.people[i];
                p.boarded = true;
            }

            for (var i in self.people) {
                const p = self.people[i];
                if (!p.locked) {
                    moved = true;
                    p.y -= 1;

                    if (p.y === -5) {
                        p.locked = true;
                    } else if (p.y < -1) {
                        p.locked = Math.random() > 0.5;
                    }
                }
            }

            if (!moved) {
                clearInterval(intId);
                window.state = window.StateBoardingFinished
            }

        }, window.refreshRate);
    }

    addArrivals() {

        this.staticPeople = [];

        let count = 20;
        while (count > 0) {

            let x = Math.round(Math.random() * (this.width - 3)) + 1;
            let y = Math.round(Math.random() * (this.height - 3)) + 1;

            let person = new Person(x, y);
            person.trainStatic = true;
            person.locked = true;
            person.boarded = true;
            this.staticPeople.push(person);

            count--;
        }


    }

}