# Authentication
- Needs Proof of an identity
- Proof for application is password.
    - It is used to prove that the username that I entered is really me.

So Showing ID to hotel receptionist is a form of authentication.

# Authorization
After receptionist made sure who I am, I am handed a hotel key. The hotel key only gives me access to certain parts of the building. I can ofcourse open the door of my room, but also the door to the pool but not the doors of the rooms to other hotel guests, if only to avoid unpleasant surprises.

So the key offers us limited access to, in this case the hotel building and that is called authorization. It's about what somebody can do and we often want to limit what the user can do in an application too.

For Authorization, data is needed, and the claims that were obtained during the authentication process could be a source for that, claims like address and birth date, but claims could also be organization specific like role, employee number or department. With that information we can perform authorization.
E.g. Only if the role claim has the value manager, I can add a new customer, or only if my department claim has the value of HR, I can see the salaries of the employees. Authorization can be done by using claims directly or indirectly.

# Don't Trust The Browser
Someone came up with this idea to approach authentication. Don't do it in this way. This is just a bad example that will motivate us to do it right. The idea is to keep a list of users with their password in a React Application using an in-memory object and then create a logging component that asks for a username and password and check these against the object. And when everything checked out we know users are authorized and we can limit UI access based on their role, which also could be contained in that object.

When it's time to contact an API it would then add a HTTP header to each request with a password, also known as an API key, so that the API is protected. The API will do a check for the API key everytime a request is done. When the API key isn't present, access to it is denied.

Here is why you should never do this:
- Everything that is sent to the browser is readable and can be manipulated. So the secrets like the API key and user credential are easily obtained and misused.
- There is major problem with single-page applications like our react applications, we can't keep secrets like passwords, tokens and client secrets anywhere in the application. 
- You might think that we can encrypt things. Yes, but the problem with encryption is that a key is needed to decrypt and where do we keep that key? Only place to keep it is within the browser context somewhere and that's not safe.
- ANother problem with this approach is that access is only limited on the UI side. So we can hide the Add button for a customer when the user is not admin for example. But everything is changeable in the browser. By some simple manipulation a tech savvy user can change the code so that the button appear and add a customer anyway. And since there is no authorization rules checks outside of the API key on the API, that will just work.
    - Part of the problem is that API doesn't know who the user is, so it can't restrict access.

Now that we know the problems lets dive into real solutions.

# Backend For Frontend (BFF)
Since browsers can't keep secrets, the bottom line is that for security, we always need to rely on a server-side companion application, a backend that can keep secrets.

Only way to get authentication right and secure is to delegate the task to a server application. The pattern used to make that work is called BFF, or Backend for Frontend.

# Application Architecture
Here is the concept for the application where we'll see the code for in a minute. Auntentication aside, the fact is that any single page application, let's call it SPA from now on, needs some kind of server, atleast to be able to download the static assets. The download would contain all HTML, CSS and JavaScript needed for the application. These static assets could come from a source like a CDN. But for now let's assume the backend application hosts them.

So the first request of the browser to the backend will bring the static assets over. And after that the React application is running in the browser, and it typically does request to an API with data endpoints to get and manipulate data.

## Logging In 
How do we add authentication?
Since React application can't do it by itself we will add a **login** endpoint on the backend. It's one of the user endpoints.

So now React application can redirect to a login page rendered not using React, but on the backend application, which will send it to the browser. The user can fill out username and password, which are submitted as a form post request to the backend and there the credentials are checked. 

And if everything is okay, and this is important part, the backend sends a cookie to the browser, which is called the identity cookie because it contains the user identity in the form of claims.

That cookie will be stored in the browser and as soon as a request such as a request to the API endpoints is sent to the same domain the cookie came from, the cookie is sent along automatically. That's just standard browser behavior.

So now that we have the cookie, we can protect endpoints on the server application. For example we can configure the endpoints that fetch data to require the cookie. On each request, the cookie will be sent along and the backend will know who users are and can check their claims too, if needed.

But we are still missing one important part of the puzzle here because the problem to begin with was that everything sent to the browser is readable. So aren't we just sending sensitive information to the browser in the form of the cookie now?
That's true but what we didn't mention yet is the fact that the cookie is **encrypted** by the backend application and the key used to encrypt and decrypt is not sent to the browser. It's kept safely on the server. So the cookie's information can't be read by the browser. And it is also made temperproof by the server because it is digitally signed as well.

So now when the backend receives a request, it knows who the user is. But how can the React application know the claims of a user. After all it can't read the cookie because it's encrypted.
For that we can simply add another endpoint to the backend application that gets back the user claims.

And ofcourse that endpoint requires a valid identity cookie too.

### Cookie Hijacking
Now what if somehow the cookie is obtained by attackers?
They could still impersonate users and access endpoints on behalf of them. Now that's true but the cookie is way harder to read or obtain from JavaScript, assuming a cross-site scripting attack. First of all identity cookies have the **HttpOnly** flag, that means they can't be accessed from JavaScript at all and on top of that there's the mentioned encryption.