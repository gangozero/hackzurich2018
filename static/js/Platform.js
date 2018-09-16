class Platform {

    constructor(size) {
        this.size = size;
        this.people = [];
    }

    setWalk(distribution) {
        for (const i in this.people) {
            const person = this.people[i];
            person.walkAround(distribution, this);
        }
    }

    addPeople(n) {

        this.people = this.people.filter(function (el) {
            return (!el.boarded);
        });

        n = n / 4;
        const zoneA = n / 3;
        const zoneB = n * 4;
        const zoneC = n / 3;

        const borderA = Math.floor(this.size.width / 3);
        const borderB = borderA * 2;
        const borderC = borderA * 3;

        this.fillZone(zoneA, 0, borderA);
        this.fillZone(zoneB, borderA, borderB);
        this.fillZone(zoneC, borderB, borderC);

        this.makeVisible();
    }

    informCentralUsers(users) {

        let size = users.length - 1;
        const self = this;
        const intId = setInterval(function () {

            var count = 5;

            while (size > 0 && count > 0) {

                let u = users[size];

                for (let p of self.people) {
                    if (p.id === u.user_id) {
                        p.custom = true;
                        p.wentToFirstClass = u.minor === 31;
                        p.distributionLocation = u.minor === 31 ? self.getRightRandomX() : self.getLeftRandomX();
                        p.initX = p.distributionLocation;
                    }
                }
                size--;
                count--;
            }

            if (size === 0) {
                window.state = window.StateStartDistribution;
                clearInterval(intId);
            }

        }, window.refreshRate);
    }

    getLeftRandomX() {
        let part = this.size.width / 3 - 2;
        let result = Math.round(Math.random() * part) + 1;
        return result;
    }

    getRightRandomX() {
        let part = this.size.width / 3;
        let result = Math.round(part * 2 + Math.random() * part) - 2;
        return result;
    }

    makeVisible() {
        const self = this;
        const intId = setInterval(function () {

            let allVisible = true;
            let showCount = 20;
            for (const i in self.people) {
                const p = self.people[i];
                if (!p.visible) {
                    p.visible = true;
                    allVisible = false;
                    showCount -= 1;
                }

                if (showCount === 0) {
                    break;
                }
            }

            if (allVisible) {
                clearInterval(intId);
                window.state = window.StateWaiting;
            }

        }, window.refreshRate);
    }

    fillZone(n, left, right) {

        const width = right - left;

        while (n > 0) {

            const x = left + Math.floor(Math.random() * width);
            const y = Math.floor(Math.random() * this.size.height);

            if (
                x !== 0 && y !== 0 &&
                x !== this.size.width - 1 && y !== this.size.height - 1
            ) {
                let person = new Person(x, y);
                this.people.push(person);
                n--;
            }
        }
    }
}
