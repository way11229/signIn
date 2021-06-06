
import axios from 'axios';
import lockr from 'lockr';

require('../scss/index.scss');

const BtnLine = document.getElementById('btn-line');

if (BtnLine) {
    BtnLine.onclick = function (this: GlobalEventHandlers, ev: MouseEvent): any {
        const state: string = 'line_' + Date.now();

        lockr.set('state', state);

        location.href = `https://access.line.me/oauth2/v2.1/authorize?response_type=code&client_id=${process.env.LINE_CLIENT_ID}&redirect_uri=${process.env.LINE_REDIECT_URL}&state=${state}&scope=profile%20openid%20email`;
    }
}

window.onload = (ev: Event): void => {
    const queyString: string = window.location.search;
    if (queyString === '') {
        return;
    }

    const urlParams: URLSearchParams = new URLSearchParams(queyString);
    const state: string = lockr.get('state', '');

    if (state === '') {
        return;
    }

    if (state.includes('line')) {
        if (urlParams.has('error')) {
            alert(urlParams.get('error'));
        } else if (urlParams.has('code') && urlParams.has('state')) {
            if (urlParams.get('state') !== state) {
                alert('State error');
                return;
            }

            axios.post(
                `${process.env.API_BASE_URL}line`,
                {
                    method: 'line',
                    verifyCode: urlParams.get('code'),
                },
                {
                    headers: {
                        'Content-Type': 'application/x-www-form-urlencoded;charset=utf-8'
                    }
                }
            ).then((response: object) => {
                console.log(response);
            });
        }
    }
};