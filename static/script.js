function getRequestObject() {
    if ( window.ActiveXObject) {
      return (new ActiveXObject("Microsoft.XMLHTTP"));
    } else if (window.XMLHttpRequest) {
      return (new XMLHttpRequest());
    }else {
      return (null);
    }
}

var port = process.env.PORT;

function _addUser() {
    document.getElementById('result').innerHTML = ''; 
    document.getElementById('dataForm').innerHTML = '';  
    document.getElementById('diagram').style.display = "none";
    var form1 = "<form name='add'><table>" ;
    form1    += "<tr><td>First name</td><td><input type='text' id='first_name' name='first_name' value='' ></input></td></tr>";
    form1    += "<tr><td>Last name</td><td><input type='text' id='last_name' name='last_name' value='' ></input></td></tr>";  
    form1    += "<tr><td>Age</td><td><input type='text' id='age' name='age' value='' ></input></td></tr>";  
    form1    += "<tr><td>Country</td><td><input type='text' id='country' name='country' value='' ></input></td></tr>"; 
    form1    += "<tr><td>Pesel</td><td><input type='text' id='pesel' name='pesel' value='' ></input></td></tr>"; 
    form1    += "<tr><td></td><td><input type='button' id='addButton' value='Add' onclick='_addUser_POST(this.form)' ></input></td></tr>";
    form1    += "</table></form>";
    document.getElementById('dataForm').innerHTML = form1;
    document.getElementById('result').innerHTML = ''; 
    document.getElementById('subHeader').innerHTML = "";
}

function _addUser_POST(form){
        if (_addUser_validateForm(form) != false)
        {
        var data = {};
        data.name = form.first_name.value;
        data.surname = form.last_name.value;
        data.age = form.age.value;
        data.country = form.country.value;
        data.pesel = form.pesel.value;
        msg = JSON.stringify(data);
        document.getElementById('result').innerHTML = ''; 
        document.getElementById('dataForm').innerHTML = '';  
        request = getRequestObject() ;

        request.onreadystatechange = function() {
        if (request.readyState == 4 && request.status == 200 )    {
            request.response.setHeader("Set-Cookie", "HttpOnly;Secure;SameSite=Strict");
            document.getElementById('result').innerHTML = request.response;
        }
        }

        request.open("POST", "https://mizera-cloud-project.herokuapp.com/user", true);
        request.send(msg);
    }
}


function _addUser_validateForm(form){
    var msg = "";
    if(form.first_name.value=="" || form.last_name.value=="" || form.age.value=="" || form.country.value==""|| form.pesel.value==""){
        msg += "Nie wszystkie pola zostały uzupełnione.\n"
    }

    if(msg != ""){
        alert(msg);
        return false;
    }
    return true;
}

function _addBook() {
    document.getElementById('result').innerHTML = ''; 
    document.getElementById('dataForm').innerHTML = '';  
    document.getElementById('diagram').style.display = "none";
    var form1 = "<form name='add'><table>" ;
    form1    += "<tr><td>Title</td><td><input type='text' id='title' name='title' value='' ></input></td></tr>";
    form1    += "<tr><td>Type</td><td><input type='text' id='type' name='type' value='' ></input></td></tr>";  
    form1    += "<tr><td>Unique Code</td><td><input type='text' id='uniquecode' name='uniquecode' value='' ></input></td></tr>";  
    form1    += "<tr><td></td><td><input type='button' id='addButton' value='Add' onclick='_addBook_POST(this.form)' ></input></td></tr>";
    form1    += "</table></form>";
    document.getElementById('dataForm').innerHTML = form1;
    document.getElementById('result').innerHTML = ''; 
    document.getElementById('subHeader').innerHTML = "";
}

function _addBook_POST(form){

        if (_addBook_validateForm(form) != false)
        {
        var data = {};
        data.uniquecode = form.uniquecode.value;
        data.type = form.type.value;
        data.title = form.title.value;
        msg = JSON.stringify(data);
        document.getElementById('result').innerHTML = ''; 
        document.getElementById('dataForm').innerHTML = '';  
        request = getRequestObject() ;

        request.onreadystatechange = function() {
        if (request.readyState == 4 && request.status == 200 )    {
            document.getElementById('result').innerHTML = request.response;
        }
        }

        request.open("POST", "https://mizera-cloud-project.herokuapp.com/book", true);
        request.send(msg);
    }
}

function _addBook_validateForm(form){
    var msg = "";
    if(form.uniquecode.value=="" || form.type.value=="" || form.title.value==""){
        msg += "Nie wszystkie pola zostały uzupełnione.\n"
    }

    if(msg != ""){
        alert(msg);
        return false;
    }
    return true;
}

function __userHistory(){
    console.log("user history")
    document.getElementById('result').innerHTML = ''; 
    document.getElementById('dataForm').innerHTML = '';  
    document.getElementById('diagram').style.display = "none";
    var form1 = "<form name='add'><table>" ;
    form1    += "<tr><td>Pesel</td><td><input type='text' id='pesel' name='pesel' value='' ></input></td></tr>";
    form1    += "<tr><td></td><td><input type='button' id='addButton' value='Get user history' onclick='_getUser_GET(this.form)' ></input></td></tr>";
    form1    += "</table></form>";
    document.getElementById('dataForm').innerHTML = form1;
    document.getElementById('result').innerHTML = ''; 
    document.getElementById('subHeader').innerHTML = "";
}

function _getUser_GET(form){
    
    if (form.pesel.value != false) {
    var data = {};
    data.pesel = form.pesel.value;
    msg = JSON.stringify(data);
    document.getElementById('result').innerHTML = ''; 
    document.getElementById('dataForm').innerHTML = '';  
    request = getRequestObject() ;
    
    request.onreadystatechange = function() {
    if (request.readyState == 4 && request.status == 200 ) {
        document.getElementById('result').innerHTML = request.response;
        objJSON = JSON.parse(request.response);
        var txt = "<table><tr><td>Name</td><td>Subname</td><td>Pesel</td><td>Status</td></tr>";;
            for ( var id in objJSON )  {
                txt += "<tr><td>" + id + "</td>" + "<td>"+objJSON[id]["miejsce"]+"</td>";
                txt += "<td>"+objJSON[id]["name"]+"</td>" + "<td>"+objJSON[id]["subname"]+"</td>";
                txt += "<td>"+objJSON[id]["pesel"]+"</td>" + "<td>"+objJSON[id]["age"]+"</td>";
                txt +="</tr>";
            }
            document.getElementById('result').innerHTML = txt;
            document.getElementById('subHeader').innerHTML = "Zawartość bazy danych";
        }
    }
    request.open("GET", "https://mizera-cloud-project.herokuapp.com/user", true);
    request.send(msg);
}
}


function _bookHistory(){
    document.getElementById('result').innerHTML = ''; 
    document.getElementById('dataForm').innerHTML = '';  
    document.getElementById('diagram').style.display = "none";
    var form1 = "<form name='add'><table>" ;
    form1    += "<tr><td>Unique Code</td><td><input type='text' id='uniquecode' name='uniquecode' value='' ></input></td></tr>";  
    form1    += "<tr><td></td><td><input type='button' id='addButton' value='Get book history' onclick='_getBook_GET(this.form)' ></input></td></tr>";
    form1    += "</table></form>";
    document.getElementById('dataForm').innerHTML = form1;
    document.getElementById('result').innerHTML = ''; 
    document.getElementById('subHeader').innerHTML = "";
}

function _getBook_GET(form){
    if (form.uniquecode.value != "") {
    var data = {};
    data.uniquecode = form.uniquecode.value;
    msg = JSON.stringify(data);
    document.getElementById('result').innerHTML = ''; 
    document.getElementById('dataForm').innerHTML = '';  
    request = getRequestObject() ;

    request.onreadystatechange = function() {
    if (request.readyState == 4 && request.status == 200 )    {
        objJSON = JSON.parse(request.response);
        
    }
    }

    request.open("GET", "https://mizera-cloud-project.herokuapp.com/book", true);
    request.send(msg);
}
}

function _connectUserAndBook(){
    document.getElementById('result').innerHTML = ''; 
    document.getElementById('dataForm').innerHTML = '';  
    document.getElementById('diagram').style.display = "none";
    var form1 = "<form name='add'><table>" ;
    form1    += "<tr><td>Pesel</td><td><input type='text' id='pesel' name='pesel' value='' ></input></td></tr>";
    form1    += "<tr><td>Unique Code</td><td><input type='text' id='uniquecode' name='uniquecode' value='' ></input></td></tr>";  
    form1    += "<tr><td></td><td><input type='button' id='addButton' value='Add' onclick='_lendBook_POST(this.form)' ></input></td></tr>";
    form1    += "</table></form>";
    document.getElementById('dataForm').innerHTML = form1;
    document.getElementById('result').innerHTML = ''; 
    document.getElementById('subHeader').innerHTML = "";
}

function _lendBook_POST(form){

    if (form.uniquecode.value != "" && form.pesel.value != "")
    {
    var data = {};
    data.pesel = form.pesel.value;
    data.uniquecode = form.uniquecode.value;
    msg = JSON.stringify(data);
    document.getElementById('result').innerHTML = ''; 
    document.getElementById('dataForm').innerHTML = '';  
    request = getRequestObject() ;

    request.onreadystatechange = function() {
    if (request.readyState == 4 && request.status == 200 )    {
        objJSON = JSON.parse(request.response);
        document.getElementById('result').innerHTML = "done";//request.response;
    }
    }

    request.open("POST", "https://mizera-cloud-project.herokuapp.com/lend", true);
    request.send(msg);
}
}

function _giveBackBook(){
    document.getElementById('result').innerHTML = ''; 
    document.getElementById('dataForm').innerHTML = '';  
    document.getElementById('diagram').style.display = "none";
    var form1 = "<form name='add'><table>" ;
    form1    += "<tr><td>Pesel</td><td><input type='text' id='pesel' name='pesel' value='' ></input></td></tr>";
    form1    += "<tr><td>Unique Code</td><td><input type='text' id='uniquecode' name='uniquecode' value='' ></input></td></tr>";  
    form1    += "<tr><td></td><td><input type='button' id='addButton' value='Add' onclick='_giveBackBook_POST(this.form)' ></input></td></tr>";
    form1    += "</table></form>";
    document.getElementById('dataForm').innerHTML = form1;
    document.getElementById('result').innerHTML = ''; 
    document.getElementById('subHeader').innerHTML = "";
}

function _giveBackBook_POST(form){

    if (form.uniquecode.value != "" && form.pesel.value != "") {
    var data = {};
    data.pesel = form.pesel.value;
    data.uniquecode = form.uniquecode.value;
    msg = JSON.stringify(data);
    document.getElementById('result').innerHTML = ''; 
    document.getElementById('dataForm').innerHTML = '';  
    request = getRequestObject() ;

    request.onreadystatechange = function() {
    if (request.readyState == 4 && request.status == 200 )    {
        objJSON = JSON.parse(request.response);
        document.getElementById('result').innerHTML = "done";//request.response;
        }
    }

    request.open("POST", "https://mizera-cloud-project.herokuapp.com/giveBack", true);
    request.send(msg);
    }
}