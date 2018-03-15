// document.getElementById("btn-start").onclick = function () {
//     location.href = "/";
// };
let startButton = document.getElementById('btn-start');
let startMenu = document.getElementById('os-app-sm');

startButton.onclick = function () {
    if (startMenu.style.display === "block") {
        startMenu.style.display = "none";
    } else {
        startMenu.style.display = "block";
    }
};

document.addEventListener('DOMContentLoaded', function() {
    document.getElementById("btn-start-exit").addEventListener("click", home);
    document.getElementById("os-app-btn-home").addEventListener("click", home);

    document.getElementById("front-nav-btn-applications").addEventListener("click", osModApplications);
    document.getElementById("front-nav-btn-business").addEventListener("click", osModBusiness);
    document.getElementById("front-nav-btn-services").addEventListener("click", osModServices);
    document.getElementById("front-nav-btn-products").addEventListener("click", osModProducts);
    document.getElementById("front-nav-btn-promotions").addEventListener("click", osModPromotions);

    document.getElementById("os-mod-link-about").addEventListener("click", osAbout);
    document.getElementById("os-mod-link-test").addEventListener("click", osTest);

    document.getElementById("start-menu-btn-open-crm").addEventListener("click", appCrm);
    document.getElementById("start-menu-btn-open-crm-dashboard").addEventListener("click", appCrmDashboard);
    document.getElementById("start-menu-btn-open-crm-customers").addEventListener("click", appCrmCustomers);
});

function home() {
    location.href = "/";
}

function osAbout() {
    location.href = "/about";
}
function osTest() {
    location.href = "/test";
}

function osModApplications() {
    location.href = "/apps";
}
function osModBusiness() {
    location.href = "/business";
}
function osModServices() {
    location.href = "/services";
}
function osModProducts() {
    location.href = "/products";
}
function osModPromotions() {
    location.href = "/promotions";
}

function appCrm() {
    location.href = "/apps/crm";
}
function appCrmDashboard() {
    location.href = "/apps/crm/dashboard";
}
function appCrmCustomers() {
    location.href = "/apps/crm/customers";
}