//https://stackoverflow.com/questions/49244944/toggle-appendchild-if-doesnt-exist-and-removechild-if-exist/49245337#49245337
//Will play with this to use display: block and outside click to hide.
function toggle() {
    'use strict';

    let element = document.getElementById('os-start-menu');
    let content = element.content;
    
    // Attempt to reference the element in the document, not the template content
    var imported = document.querySelector(".os1-start-menu");

    // Check for the element, not the template content
    if (document.body.contains(imported)) {
        // Element exists, get its parent
        imported.parentNode.removeChild(imported);
    } else {
        // Use .importNode to bring template content in:
        document.body.appendChild(document.importNode(content, true));
    }
}

let btnOpenLauncher = document.getElementById('btn-open-launcher');
if (btnOpenLauncher) {
    btnOpenLauncher.addEventListener('click', toggle, false);
}

let btnCloseApp = document.getElementById("btn-close-app");
if (btnCloseApp) {
    btnCloseApp.addEventListener("click", goToDesktop, false);
}

let signOut = document.getElementById("btn-sign-out");
if (signOut) {
    signOut.addEventListener("click", goToMain, false);
}

function goToDesktop() {
    location.href = "/desktop";
};
function goToMain() {
    location.href = "/";
};

var mymap = L.map('mapid').setView([56.94965, 24.1052], 13);
var marker = L.marker([56.94965, 24.1052]).addTo(mymap);

L.tileLayer('https://api.tiles.mapbox.com/v4/{id}/{z}/{x}/{y}.png?access_token={accessToken}', {
    attribution: 'Map data &copy; <a href="https://www.openstreetmap.org/">OpenStreetMap</a> contributors, <a href="https://creativecommons.org/licenses/by-sa/2.0/">CC-BY-SA</a>, Imagery © <a href="https://www.mapbox.com/">Mapbox</a>',
    maxZoom: 16,
    id: 'mapbox.streets',
    accessToken: 'pk.eyJ1Ijoib3N3ZWUiLCJhIjoiY2ppMXdycTZyMGY3ZDNwcXBzZnUzcTBiYiJ9.mi4BGXgJbdDgS7mVGVyExA',
    style: 'mapbox://styles/mapbox/basic-v9',
}).addTo(mymap);