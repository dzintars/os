(function() {
    /**
     * Define key codes to help with handling keyboard events.
     */
    const KEYCODE = {
        SPACE: 32,
        ENTER: 13,
    };
    
    /**
     * Cloning contents from a &lt;template&gt; element is more performant
     * than using innerHTML because it avoids additional HTML parse costs.
     */
    
    const template = document.createElement('template');
  
     template.innerHTML = `
     <style>
        :host {
            display: flex;
            flex-direction: row;
            border: none;
            background: #006AFF;
            padding: 3px;
            width: 80px;
        }
        :host([pressed]) {
            background: lightgreen;
        }
        :host([disabled]) {
            background: lightgray;
        }
        </style>
        <os-dialog></os-dialog>
        <slot></slot>
     `;
    
     class OsButtonToggle extends HTMLElement {
        static get observedAttributes() {
            return ['pressed','disabled'];
        }
    
        constructor() {
            super();
            this.attachShadow({mode: 'open'});
            this.shadowRoot.appendChild(template.content.cloneNode(true));
        }
    
        connectedCallback() {
            if (!this.hasAttribute('role')) {
                this.setAttribute('role', 'button');
            }
            if (!this.hasAttribute('tabindex')) {
                this.setAttribute('tabindex', '0');
            }
            if (!this.hasAttribute('aria-pressed')) {
                this.setAttribute('aria-pressed', 'false');
            }
    
            this.addEventListener('click', this._onClick);
            this.addEventListener('keydown', this._onKeyDown);
        }
    
        set pressed(value) {
            const isPressed = Boolean(value);
            if (isPressed) {
                this.setAttribute('pressed', '');
            } else {
                this.removeAttribute('pressed');
            }
        }
        get pressed() {
            return this.hasAttribute('pressed');
        }
    
        set disabled(value) {
            const isDisabled = Boolean(value);
            if (isDisabled) {
                this.setAttribute('disabled', '');
            } else {
                this.removeAttribute('disabled');
            }
        }
        get disabled() {
            return this.hasAttribute('disabled');
        }
    
        attributeChangedCallback(name, oldVal, newVal) {
            const hasValue = newVal !== null;
            this.setAttribute(`aria-${name}`, hasValue);
        }
        _onClick() {
            this._togglePressed();
        }
    
        _onKeyDown(event) {
            if (event.altKey) {
                return;
            }
    
            switch (event.keyCode) {
                case KEYCODE.ENTER:
                case KEYCODE.SPACE:
                    event.preventDefault();
                    this,this._togglePressed();
                    break;
    
                default:
                    break;
    
            }
        }
    
        _togglePressed() {
            this.pressed = !this.pressed;
        }
     }
    
     window.customElements.define('os-button-toggle', OsButtonToggle);
})();