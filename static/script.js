var serverPath = "https://hobby-oghlklmkakojgbkepalfbfdl.dbs.graphenedb.com:24780/db/data/"// "http://localhost:8080";

function getRequestObject() {
    if ( window.ActiveXObject) {
      return (new ActiveXObject("Microsoft.XMLHTTP"));
    } else if (window.XMLHttpRequest) {
      return (new XMLHttpRequest());
    }else {
      return (null);
    }
}

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
        if (request.readyState == 4 && request.status == 201 )    {
            document.getElementById('subHeader').innerHTML = "Successfuly added user: " + data.pesel;
            }
        }

        request.open("POST", serverPath + "/user", true);
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
        if (request.readyState == 4 && request.status == 201 )    {
            document.getElementById('subHeader').innerHTML = "Successfuly added book: " + data.uniquecode;
        }
        }

        request.open("POST", serverPath + "/book", true);
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
    
    if (form.pesel.value != "") {
    var data = {};
    data.Pesel = form.pesel.value;
    document.getElementById('result').innerHTML = ''; 
    document.getElementById('dataForm').innerHTML = '';  
    request = getRequestObject() ;
    request.onreadystatechange = function() {
    if (request.readyState == 4 && request.status == 200 ) {
        objJSON = JSON.parse(request.response);
        var txt = "<table><tr><td>Title</td><td>Type</td><td>UniqueCode</td><td>Status</td></tr>";;
            for ( var id in objJSON )  {
                txt += "<tr><td>"+objJSON[id]["Title"]+"</td>" + "<td>"+objJSON[id]["Type"]+"</td>";
                txt += "<td>"+objJSON[id]["UniqueCode"]+"</td>" + "<td>"+objJSON[id]["State"]+"</td>";
                txt +="</tr>";
            }
            document.getElementById('result').innerHTML = txt + "</table>";
            document.getElementById('subHeader').innerHTML = "Account history: user " + data.Pesel;
        } else {
            document.getElementById('subHeader').innerHTML = "Account for user " + data.Pesel + " doesn't exist.";
        }
    }
    request.open("GET", serverPath + "/user/" + data.Pesel , true) ;
    request.send(null);
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
   
    document.getElementById('result').innerHTML = ''; 
    document.getElementById('dataForm').innerHTML = '';  
    request = getRequestObject() ;

    request.onreadystatechange = function() {
        if (request.readyState == 4 && request.status == 200 ) {
            objJSON = JSON.parse(request.response);
            var txt = "<table><tr><td>Name</td><td>Surname</td><td>Pesel</td><td>Status</td></tr>";;
                for ( var id in objJSON )  {
                    txt += "<tr><td>"+objJSON[id]["Name"]+"</td>" + "<td>"+objJSON[id]["Surname"]+"</td>";
                    txt += "<td>"+objJSON[id]["Pesel"]+"</td>" + "<td>"+objJSON[id]["State"]+"</td>";
                    txt +="</tr>";
                }
                document.getElementById('result').innerHTML = txt + "</table>";
                document.getElementById('subHeader').innerHTML = "Account history: book " + data.uniquecode;
        } else {
                document.getElementById('subHeader').innerHTML = "Account for book " + data.uniquecode + " doesn't exist.";
        }
    }

    request.open("GET", serverPath + "/book/"+ data.uniquecode , true);
    request.send(null);
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
    if (request.readyState == 4 && request.status == 201 ) {
        alert("Success.")
    }
    }

    request.open("POST", serverPath + "/lend", true);
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
    if (request.readyState == 4 && request.status == 201 )    {
        objJSON = JSON.parse(request.response);
        alert("Success.")
        document.getElementById('result').innerHTML = "Succesfully returned a book: " + data.uniquecode + "by user: " + data.pesel ;//request.response;
        }
    }

    request.open("POST", serverPath + "/giveBack", true);
    request.send(msg);
    }
}


function _deleteUser(){
    document.getElementById('result').innerHTML = ''; 
    document.getElementById('dataForm').innerHTML = '';  
    document.getElementById('diagram').style.display = "none";
    var form1 = "<form name='add'><table>" ;
    form1    += "<tr><td>Pesel</td><td><input type='text' id='pesel' name='pesel' value='' ></input></td></tr>";  
    form1    += "<tr><td></td><td><input type='button' id='addButton' value='Delete user' onclick='_deleteUser_DELETE(this.form)' ></input></td></tr>";
    form1    += "</table></form>";
    document.getElementById('dataForm').innerHTML = form1;
    document.getElementById('result').innerHTML = ''; 
    document.getElementById('subHeader').innerHTML = "";
}

function _deleteUser_DELETE(form){
    if (form.pesel.value != "") {
    var data = {};
    data.pesel = form.pesel.value;
   
    document.getElementById('result').innerHTML = ''; 
    document.getElementById('dataForm').innerHTML = '';  
    request = getRequestObject() ;

    request.onreadystatechange = function() {
    if (request.readyState == 4 && request.status == 204 )    {
        alert("Success")
    } else {
        alert(request.response)
    }
    }

    request.open("DELETE", serverPath + "/user/"+ data.pesel , true);
    request.send(null);
}
}
