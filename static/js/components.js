// Create a class for the element
class OsInfo extends HTMLElement {
  constructor() {
    // Always call super first in constructor
    super();

    // Create a shadow root
    let shadow = this.attachShadow({mode: 'open'});

    // Create spans
    let wrapper = document.createElement('span');
    wrapper.setAttribute('class','wrapper');

    let icon = document.createElement('span');
    icon.setAttribute('class','icon');
    icon.setAttribute('tabindex', 0);

    let info = document.createElement('span');
    info.setAttribute('class','info');

    // Take attribute content and put it inside the info span
    let text = this.getAttribute('text');
    info.textContent = text;

    // Insert icon
    let imgUrl;
    if(this.hasAttribute('img')) {
      imgUrl = this.getAttribute('img');
    } else {
      imgUrl = 'img/default.png';
    }

    let img = document.createElement('img');
    img.src = imgUrl;
    icon.appendChild(img);

    let size = this.getAttribute('size');
    icon.style.width = size;


    // Create some CSS to apply to the shadow dom
    let style = document.createElement('style');

    style.textContent = '.wrapper {' +
                           'position: relative;' +
                        '}' +

                         '.info {' +
                            'font-size: 0.8rem;' +
                            'width: 200px;' +
                            'display: inline-block;' +
                            'border: 1px solid black;' +
                            'padding: 10px;' +
                            'background: white;' +
                            'border-radius: 2px;' +
                            'opacity: 0;' +
                            'transition: 0.6s all;' +
                            'position: absolute;' +
                            'bottom: 40px;' +
                            'left: 50px;' +
                            'z-index: 3;' +
                          '}' +

                          'img {' +
                            'width: 1.2rem' +
                          '}' +

                          '.icon:hover + .info, .icon:focus + .info {' +
                            'opacity: 1;' +
                          '}';

    // attach the created elements to the shadow dom

    shadow.appendChild(style);
    shadow.appendChild(wrapper);
    wrapper.appendChild(icon);
    wrapper.appendChild(info);
  }
}

// Define the new element
customElements.define('os-info', OsInfo);