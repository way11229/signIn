
import axios from 'axios';

require('../scss/index.scss');

const BtnLine = document.getElementById('btn-line');

if (BtnLine) {
    BtnLine.onclick = function (this: GlobalEventHandlers, ev: MouseEvent): any {
        console.log('456');
        axios.get('/test')
            .then((response: object) => {
                console.log(response);
            });
    }
}