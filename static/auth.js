function Authorization() {
    username = document.getElementById("login").value;
    password = document.getElementById("password").value;
    if ((username == "admin") && (password == "admin")) {
        window.open("index.html","_self");
    } else {
        window.open("index_noadmin.html","_self");
    }
}