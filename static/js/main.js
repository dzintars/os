// document.getElementById("btn-start").onclick = function () {
//     location.href = "/";
// };
// let startButton = document.getElementById('btn-start');
// let startMenu = document.getElementById('os-app-sm');

// startButton.onclick = function () {
//     if (startMenu.style.display === "block") {
//         startMenu.style.display = "none";
//     } else {
//         startMenu.style.display = "block";
//     }
// };

// document.addEventListener('DOMContentLoaded', function() {
//     // document.getElementById("btn-start-exit").addEventListener("click", home);
//     // document.getElementById("os-app-btn-home").addEventListener("click", home);

//     document.getElementById("btn-goto-services").addEventListener("click", osModServices);
//     document.getElementById("btn-goto-products").addEventListener("click", osModProducts);
//     document.getElementById("btn-goto-business").addEventListener("click", osModBusiness);
//     document.getElementById("btn-goto-promotions").addEventListener("click", osModPromotions);
//     document.getElementById("btn-goto-apps").addEventListener("click", osModApplications);

//     // document.getElementById("os-mod-link-about").addEventListener("click", osAbout);
//     // document.getElementById("os-mod-link-test").addEventListener("click", osTest);

//     // document.getElementById("start-menu-btn-open-crm").addEventListener("click", appCrm);
//     // document.getElementById("start-menu-btn-open-crm-dashboard").addEventListener("click", appCrmDashboard);
//     // document.getElementById("start-menu-btn-open-crm-customers").addEventListener("click", appCrmCustomers);
// });

// function home() {
//     location.href = "/";
// }

// function osAbout() {
//     location.href = "/about";
// }
// function osTest() {
//     location.href = "/test";
// }

// function osModApplications() {
//     location.href = "/apps";
// }
// function osModBusiness() {
//     location.href = "/business";
// }
// function osModServices() {
//     location.href = "/services";
// }
// function osModProducts() {
//     location.href = "/products";
// }
// function osModPromotions() {
//     location.href = "/promotions";
// }

// function appCrm() {
//     location.href = "/apps/crm";
// }
// function appCrmDashboard() {
//     location.href = "/apps/crm/dashboard";
// }
// function appCrmCustomers() {
//     location.href = "/apps/crm/customers";
// }

//https://stackoverflow.com/questions/49244944/toggle-appendchild-if-doesnt-exist-and-removechild-if-exist/49245337#49245337
//Will play with this to use display: block and outside click to hide.
let element = document.getElementById('os-start-menu');
let content = element.content;

function toggle () {
    'use strict';
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

let button = document.getElementById('btn-open-launcher');
button.addEventListener('click', toggle, false);

//document.addEventListener('DOMContentLoaded', function() {
document.addEventListener('click', function(e) {
    document.getElementById("btn-goto-services").addEventListener("click", osModServices);
    document.getElementById("btn-goto-products").addEventListener("click", osModProducts);
    document.getElementById("btn-goto-business").addEventListener("click", osModBusiness);
    document.getElementById("btn-goto-promotions").addEventListener("click", osModPromotions);
    document.getElementById("btn-goto-apps").addEventListener("click", osModApplications);
});

function osModServices() {
    location.href = "/services";
}
function osModProducts() {
    location.href = "/products";
}
function osModBusiness() {
    location.href = "/business";
}
function osModPromotions() {
    location.href = "/promotions";
}
function osModApplications() {
    location.href = "/apps";
}