(function () {
    var a = ["A Y", "B X", "C"];
    var encodeMove = function (m) {
        switch (m) {
            case "A":
            case "X":
                return 0;
            case "B":
            case "Y":
                return 1;
            case "C":
            case "Z":
                return 2;
            default:
                throw new Error();
        }
    };
    var calcResult = function (r) {
        var m = r.myMove, t = r.theirMove;
        if ((m - 1) % 3 == t) {
            return 6;
        }
        else if ((m + 1) % 3 == t) {
            return 0;
        }
        else if (m == t) {
            return 3;
        }
        else {
            throw new Error();
        }
    };
    var b = a.map(function (b) {
        var me = b[2];
        var them = b[0];
        var round = {
            myMove: encodeMove(me),
            theirMove: encodeMove(them)
        };
        return round.myMove + 1 + calcResult(round);
    });
    console.log(b, b.reduce(function (a, b) { return a + b; }, 0));
})();
