
var password="admin", username="admin";

function getDrivers() {
    req = document.getElementById("id1").value;
    var xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            resp = JSON.parse(xhttp.responseText)
            document.getElementById("demo").innerHTML = resp.name.toString();
        }
    };
    url = "http://localhost:1323/drivers/route/"+req;
    xhttp.open("GET", url, true);
    xhttp.send();
}

function getBuses() {
    req = document.getElementById("id2").value;
    var xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            resp = JSON.parse(xhttp.responseText)
            document.getElementById("demo").innerHTML = resp.numbers.toString();
        }
    };
    url = "http://localhost:1323/buses/"+req;
    xhttp.open("GET", url, true);
    xhttp.send();
}
function getLengthRoutes() {
    req = document.getElementById("id3").value;
    var xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            resp = JSON.parse(xhttp.responseText)
            var route = "";
            for (i=0;i<resp.data.length;i++) {
                route += resp.data[i].route+" "+resp.data[i].length + "\n";
            }
            document.getElementById("demo").innerHTML = route;
        }
    };
    url = "http://localhost:1323/routes/"+req+"/length";
    xhttp.open("GET", url, true);
    xhttp.send();
}
function getTimeRoutes() {
    req = document.getElementById("id4").value;
    time = document.getElementById("s_id4").value;
    var xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            resp = JSON.parse(xhttp.responseText)
            var route = "";
            for (i=0;i<resp.data.length;i++) {
                route += resp.data[i].route+" "+resp.data[i].moveTime + "\n";
            }
            document.getElementById("demo").innerHTML = route;
        }
    };
    url = "http://localhost:1323/route/"+req+"/"+time;
    xhttp.open("GET", url, true);
    xhttp.send();
}
function getDriverBus() {
    req = document.getElementById("id5").value;
    var xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            resp = JSON.parse(xhttp.responseText)
            document.getElementById("demo").innerHTML = resp.name.toString();
        }
    };
    url = "http://localhost:1323/drivers/bus/"+req;
    xhttp.open("GET", url, true);
    xhttp.send();
}
function AddDriver() {
    fullname = document.getElementById("id6").value;
    cl = document.getElementById("id6_1").value;
    exp = document.getElementById("id6_2").value;
    bus = document.getElementById("id6_3").value;

    var xhttp = new XMLHttpRequest();

    xhttp.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            resp = JSON.parse(xhttp.responseText)
            document.getElementById("demo").innerHTML = resp;
        }
    };
    var Driver = JSON.stringify({ fullname: fullname, class: Number(cl), experience: Number(exp), idBus: Number(bus) });
    url = "http://localhost:1323/admin/driver";
    xhttp.open("POST", url, true);
    xhttp.setRequestHeader('Content-Type', 'application/json');
    xhttp.setRequestHeader("Authorization", "Basic " + btoa(username + ":" + password));
    xhttp.send(Driver);
}
function PutLength() {
    route = document.getElementById("id7").value;
    length = document.getElementById("id7_1").value;

    var xhttp = new XMLHttpRequest();

    xhttp.onreadystatechange = function() {

        if (this.readyState == 4 && this.status == 200) {
            resp = JSON.parse(xhttp.responseText)
            document.getElementById("demo").innerHTML = resp;
        }
    };
    var Driver = JSON.stringify({ route: Number(route), length: Number(length)});
    url = "http://localhost:1323/admin/bus/length";
    xhttp.open("PUT", url, true);
    xhttp.setRequestHeader('Content-Type', 'application/json');
    xhttp.setRequestHeader("Authorization", "Basic " + btoa(username + ":" + password));
    xhttp.send(Driver);
}
function DeleteBus() {
    bus = document.getElementById("id8").value;

    var xhttp = new XMLHttpRequest();

    xhttp.onreadystatechange = function() {

        if (this.readyState == 4 && this.status == 200) {
            resp = JSON.parse(xhttp.responseText)
            document.getElementById("demo").innerHTML = resp;
        }
    };
    url = "http://localhost:1323/admin/bus/del/"+bus;
    xhttp.open("DELETE", url, true);
    xhttp.setRequestHeader("Authorization", "Basic " + btoa(username + ":" + password));
    xhttp.send();

}

function getInfo() {
    var xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function() {
        if (this.readyState == 4 && this.status == 200) {
            resp = JSON.parse(xhttp.responseText);
            var route = "";
            for (i=0;i<resp.routes.length;i++) {
                route += resp.routes[i].route+" "+resp.routes[i].moveStart + " "+resp.routes[i].busInterval + "\n";
            }
            var drivers = "";
            for (i=0;i<resp.drivers.length;i++) {
                drivers += resp.drivers[i].fullname+" "+resp.drivers[i].class+"\n";
            }


            document.getElementById("demo").innerHTML = "количество автобусов: " + resp.countBus + "\nтипы автобусов в парке: " + resp.busTypes.toString() + "\nномер маршрута   время начала движения   интервал\n" + route + "\nФИО   класс\n" + drivers;
        }
    };
    xhttp.open("GET", "http://localhost:1323/info", true);


    xhttp.send();
}