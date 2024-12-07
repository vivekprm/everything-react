Since browser can't keep secrets. Single Page Applications need a Server Application that can keep secrets.

The pattern used to make that work is called BFF, or Backend for Frontend.

# Application Architecture
Any single page application needs some kind of server at least to be able to download the static assets. That download would contain all HTML, CSS and Javascript needed for that application. These static assets could come from a source like CDN. But now assume backend application hosts them.

So, the First request of the browser to the backend will bring the static assets over. And after that the React Application is running in the browser and it typically does requests to an API with data endpoints to get and manipulate data.

How do we add authentication?
Since React Applications can't do it by itself, we add a login endpoint on the backend. It's one of the user endpoints. So not react application can redirect to a login page, rendered not using react, but on the backend application, which will send it to browser.
The user can fill out the username and password which are submitted as a form post request to the backend and there the credentials are checked and if everything is okay, backend sends a cookie to the browser, which is called the identity cookie because it contains the user identity in the forms of claims.
THat cookie will be stored in the browser and as soon as request such as a request to the API endpoints is sent to the same domain the cookie came from, the cookie is sent along automatically, that's just standard browser behaviour. So now we have the cookie we can protect endpoints on the server applications. For Example: We can configure the endpoints that fetch data to require the cookie. On each request the cookie will be sent along, and the backend will know who users are and can check their claims too, if needed.

But we're still missing one important part of the puzzle here because the problem to begin with was that everything sent to the browser is readable. So aren't we just sending sensitive information to the browser in the form of the cookie now?
That's true but cookie is encrypted by the backend application and the key used to encrypt and decrypt is not sent to the browser. It's kept safely on the server. So, the cookie's information can't be read by the browser and it's also made temperproof by the server because it's digitally signed as well.

But how can the React application know the claims of a user? After all it can't read the cookie because it's encrypted.
For that we can simply just add another endpoint to the backend application that gets back user claims. And ofcourse that endpoint requires a valid identity cookie too.

# Cookie Hijacking
Stealing a cookie will enable attacker to impersonate user. But the cookie is way harder to read or obtain from javascript, assuming a cross-site scripting attack.

First of all identity cookies have the **HttpOnly** flag, and that means they can't be accessed from javascript at all. And on top of that we have encryption.