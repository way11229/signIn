
import axios from 'axios';
import lockr from 'lockr';

require('../scss/index.scss');

declare global {
    interface Window { fbAsyncInit: any; }
}

const BtnLine = document.getElementById('btn-line');
const BtnFB = document.getElementById('btn-fb');

if (BtnLine) {
    BtnLine.onclick = function (this: GlobalEventHandlers, ev: MouseEvent): void {
        const state: string = 'line_' + Date.now();

        lockr.set('state', state);

        location.href = `https://access.line.me/oauth2/v2.1/authorize?response_type=code&client_id=${process.env.LINE_CLIENT_ID}&redirect_uri=${process.env.LINE_REDIECT_URL}&state=${state}&scope=profile%20openid%20email`;
    }
}

if (BtnFB) {
    BtnFB.onclick = function (this: GlobalEventHandlers, ev: MouseEvent): void {
        FB.login((response: fb.StatusResponse) => {
            if (response.status === 'connected') {
                const bodyFormData = new FormData();
                bodyFormData.append('method', 'fb');
                bodyFormData.append('verifyCode', response.authResponse.accessToken);
                bodyFormData.append('extra', JSON.stringify({
                    expireIn: response.authResponse.reauthorize_required_in,
                    userId: response.authResponse.userID,
                }));

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
            } else {
                alert('fb login error');
            }
        }, { scope: 'public_profile,email,birthday'});
    }
}

function fbInit(): void {
    window.fbAsyncInit = function () {
        FB.init({
            appId: process.env.FB_APP_ID,
            cookie: true,
            xfbml: true,
            version: 'v11.0'
        });
    };
}

function lineBasck(): void {
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

            const bodyFormData = new FormData();
            bodyFormData.append('method', 'line');
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
    }
}

window.onload = (ev: Event): void => {
    fbInit();
    lineBasck();
};