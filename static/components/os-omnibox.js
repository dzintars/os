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
    div {
        display: flex;
        align-items: center;
        border-radius: 4px;
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
    button {
        border: none;
        padding: 0 10px;
        height: 100%
    }
    button:hover {
        background-color: gray;
    }
    button:focus {
        outline: none;
    }
    </style>
    <div>
        <input type="search" name="search"></input>
        <button>Search</button>
    </div>
`;

customElements.define('os-omnibox', class extends HTMLElement {
    constructor() {
        super();
        let shadow = this.attachShadow({mode: 'open'});
        shadow.appendChild(omnibox.content.cloneNode(true))
    }
});