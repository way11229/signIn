
import axios from 'axios';
import lockr from 'lockr';

require('../scss/index.scss');

const BtnLine = document.getElementById('btn-line');
const BtnFB = document.getElementById('btn-fb');
const BtnGoogle = document.getElementById('btn-google');

if (BtnLine) {
    BtnLine.onclick = function (this: GlobalEventHandlers, ev: MouseEvent): void {
        const state: string = 'line_' + Date.now();

        lockr.set('state', state);

        location.href = `https://access.line.me/oauth2/v2.1/authorize?response_type=code&client_id=${process.env.LINE_CLIENT_ID}&redirect_uri=${process.env.LINE_REDIECT_URL}&state=${state}&scope=profile%20openid%20email`;
    }
}

if (BtnFB) {
    BtnFB.onclick = function (this: GlobalEventHandlers, ev: MouseEvent): void {
        const state: string = 'fb_' + Date.now();

        lockr.set('state', state);

        location.href = `https://www.facebook.com/v11.0/dialog/oauth?response_type=code&client_id=${process.env.FB_APP_ID}&redirect_uri=${process.env.FB_REDIECT_URL}&state=${state}&scope=public_profile,email,user_birthday`;
    }
}

if (BtnGoogle) {
    BtnGoogle.onclick = function (this: GlobalEventHandlers, ev: MouseEvent): void {
        const state: string = 'google_' + Date.now();

        lockr.set('state', state);

        location.href = `https://accounts.google.com/o/oauth2/v2/auth?response_type=code&client_id=${process.env.GOOGLE_APP_ID}&redirect_uri=${process.env.GOOGLE_REDIECT_URL}&state=${state}&scope=openid%20email%20profile`;
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

    if (urlParams.has('error')) {
        alert(urlParams.get('error'));
        return;
    }


    if (urlParams.has('code') && urlParams.has('state')) {
        if (urlParams.get('state') !== state) {
            alert('State error');
            return;
        }

        const stateSplit = state.split('_');

        const bodyFormData = new FormData();
        bodyFormData.append('method', stateSplit[0] || "");
        bodyFormData.append('verifyCode', urlParams.get('code') || "");

        axios.post(
            `${process.env.API_BASE_URL}`,
            bodyFormData,
            {
                headers: {
                    'Content-Type': 'application/x-www-form-urlencoded;charset=utf-8'
                }
            }
        ).then((response: object) => {
            console.log(response);
        });
    }
};