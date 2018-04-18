const omnibox = document.createElement('template');

omnibox.innerHTML = `
    <style>
    :host {
        box-shadow: 0 0 10px rgba(0, 0, 0, 0.3);
        height: 40px;
        width: 600px;
        max-width: 700px;
        min-width: 200px;
        margin: 20px;
    }
    form {
        display: flex;
        align-items: center;
    }
    input {
        border: none;
        height: 100%;
        width: 100%;
        padding: 0 10px;
    }
    input:focus {
        outline: none;
    }
    input[search] {
        padding: 0 10px;
        height: 100%;
        width: 80px;
    }
    input[search]:hover {
        background-color: gray;
    }
    input[search]:focus {
        outline: none;
    }
    </style>
    <form action="/search" method="GET">
        <input type="search" name="search">
        <input search type="submit" value="Search">
    </form>
`;

customElements.define('os-omnibox', class extends HTMLElement {
    constructor() {
        super();
        let shadow = this.attachShadow({mode: 'open'});
        shadow.appendChild(omnibox.content.cloneNode(true))
    }
});