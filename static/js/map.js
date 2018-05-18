var markers = [
    {{ range .Locations}}
    {
        coords:{lat:{{ .Lat}}, lng: {{ .Lng}}},
        iconImage: '{{ .Icon }}',
        content: '<p>{{ .Address }}</p>',
    },
    {{ end }}
];