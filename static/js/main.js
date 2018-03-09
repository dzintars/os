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
    document.getElementById("app-title-crm").addEventListener("click", crm);
    document.getElementById("btn-crm-customers").addEventListener("click", crmCustomers);
    document.getElementById("btn-crm-dashboard").addEventListener("click", crmDashboard);
    document.getElementById("btn-crm-projects").addEventListener("click", crmProjects);
    document.getElementById("btn-start-exit").addEventListener("click", home);
    document.getElementById("os-app-btn-home").addEventListener("click", home);
    document.getElementById("os-mod-link-about").addEventListener("click", osAbout);
    document.getElementById("os-mod-link-test").addEventListener("click", osTest);
});

function crm() {
    location.href = "/apps/crm";
}

function crmCustomers() {
    location.href = "/apps/crm/customers";
}

function crmDashboard() {
    location.href = "/apps/crm/dashboard";
}

function crmProjects() {
    location.href = "/apps/crm/projects";
}

function home() {
    location.href = "/";
}

function osAbout() {
    location.href = "/about";
}

function osTest() {
    location.href = "/test";
}