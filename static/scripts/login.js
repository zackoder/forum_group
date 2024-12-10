function validlogin(e) {

    
    e.preventDefault();
    var username = document.getElementById("email");
    var password = document.getElementById("password");
    var emailRegex = /^[a-zA-Z0-9]+@[a-zA-Z]{3,8}[.][a-zA-Z]{3,5}/;
    var isValidusrname = true;
    var isValidpassword = true;
    if (username.value === "") {
        document.getElementById("errorUsername").innerHTML = "username is null";
        document.getElementById("errorUsername").style.color = "red"
        isValidusrname=false

    } else if (username.value !== "" && !emailRegex.test(username.value)) {

        document.getElementById("errorUsername").innerHTML = "email is not valid";
        document.getElementById("errorUsername").style.color = "red"
        isValidusrname=false

    } else {
        document.getElementById("errorUsername").innerHTML = "";
        isValidusrname=true

    }

    if (password.value === "") {
        document.getElementById("errorPassword").innerHTML = "Password is null";
        document.getElementById("errorPassword").style.color = "red"
        isValidpassword=false

    } else {
        document.getElementById("errorPassword").innerHTML = "";
        isValidpassword=true
    }

    if (isValidpassword && isValidusrname) {
        document.querySelector("form").submit();
    }

}



