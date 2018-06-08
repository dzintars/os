let leafletMap = document.querySelector('#mapid');
if (leafletMap) {
    var mymap = L.map('mapid').setView([56.94965, 24.1052], 13);
    var marker = L.marker([56.94965, 24.1052]).addTo(mymap);

    L.tileLayer('https://api.tiles.mapbox.com/v4/{id}/{z}/{x}/{y}.png?access_token={accessToken}', {
        attribution: 'Map data &copy; <a href="https://www.openstreetmap.org/">OpenStreetMap</a> contributors, <a href="https://creativecommons.org/licenses/by-sa/2.0/">CC-BY-SA</a>, Imagery © <a href="https://www.mapbox.com/">Mapbox</a>',
        maxZoom: 24,
        id: 'mapbox.streets',
        accessToken: 'pk.eyJ1Ijoib3N3ZWUiLCJhIjoiY2ppMXdycTZyMGY3ZDNwcXBzZnUzcTBiYiJ9.mi4BGXgJbdDgS7mVGVyExA',
        style: 'mapbox://styles/mapbox/basic-v9',
    }).addTo(mymap);
}