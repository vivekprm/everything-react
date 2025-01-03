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

<img width="788" alt="Screenshot 2025-01-02 at 3 42 07 PM" src="https://github.com/user-attachments/assets/c276a521-d1fd-4600-9c6b-451835932f46" />

So the first request of the browser to the backend will bring the static assets over. And after that the React application is running in the browser, and it typically does request to an API with data endpoints to get and manipulate data.

<img width="568" alt="Screenshot 2025-01-02 at 3 42 26 PM" src="https://github.com/user-attachments/assets/cf7a734e-7842-4d60-abf1-d8b78087cf7d" />

## Logging In 
How do we add authentication?
Since React application can't do it by itself we will add a **login** endpoint on the backend. It's one of the user endpoints.

<img width="520" alt="Screenshot 2025-01-02 at 3 44 43 PM" src="https://github.com/user-attachments/assets/da0cdc94-0610-4c65-b81d-6c19b30909b9" />

So now React application can redirect to a login page rendered not using React, but on the backend application, which will send it to the browser. The user can fill out username and password, which are submitted as a form post request to the backend and there the credentials are checked. 

And if everything is okay, and this is important part, the backend sends a cookie to the browser, which is called the identity cookie because it contains the user identity in the form of claims.

<img width="715" alt="Screenshot 2025-01-02 at 4 26 13 PM" src="https://github.com/user-attachments/assets/37f8b860-7af1-4137-a288-c8e79b87e719" />

That cookie will be stored in the browser and as soon as a request such as a request to the API endpoints is sent to the same domain the cookie came from, the cookie is sent along automatically. That's just standard browser behavior.

<img width="683" alt="Screenshot 2025-01-02 at 4 28 18 PM" src="https://github.com/user-attachments/assets/15f59794-08ed-42be-9fa7-6e79bed8cf5c" />

So now that we have the cookie, we can protect endpoints on the server application. For example we can configure the endpoints that fetch data to require the cookie. On each request, the cookie will be sent along and the backend will know who users are and can check their claims too, if needed.

But we are still missing one important part of the puzzle here because the problem to begin with was that everything sent to the browser is readable. So aren't we just sending sensitive information to the browser in the form of the cookie now?
That's true but what we didn't mention yet is the fact that the cookie is **encrypted** by the backend application and the key used to encrypt and decrypt is not sent to the browser. It's kept safely on the server. So the cookie's information can't be read by the browser. And it is also made temperproof by the server because it is digitally signed as well.

So now when the backend receives a request, it knows who the user is. But how can the React application know the claims of a user. After all it can't read the cookie because it's encrypted.
For that we can simply add another endpoint to the backend application that gets back the user claims.

<img width="988" alt="Screenshot 2025-01-02 at 4 42 44 PM" src="https://github.com/user-attachments/assets/b3f6bb1b-e0e2-4e43-b300-eccfc029d1b4" />

And ofcourse that endpoint requires a valid identity cookie too.

### Cookie Hijacking
Now what if somehow the cookie is obtained by attackers?
They could still impersonate users and access endpoints on behalf of them. Now that's true but the cookie is way harder to read or obtain from JavaScript, assuming a cross-site scripting attack. First of all identity cookies have the **HttpOnly** flag, that means they can't be accessed from JavaScript at all and on top of that there's the mentioned encryption.

https://github.com/RolandGuijt/ps-reactauth

## Backend
Request travels through Request pipeline back and forth, reading and manipulating the request data into a response. Pieces of functionality in the pipeline are called middleware. There is middleware in the pipeline that does routing. In other words, it maps a particular URL to an action the backend application should do.
For example, a GET request on ```/customers``` could return a collection of customers, or accessing ```/account/login``` will return HTML, representing the login page.

There is middleware that enables downloading of static files such as HTML and JavaScript files and for authentication.
In our case Job of this authentication middleware is to verify, decrypt and read the identity cookie.

### Flow
We have AccountController which has number of methods called actions that have to do with user management.

Using login action login is executed when the endpoint ```/account/login``` is accessed. In that case a view is displayed with an empty model object.

When the form is submitted post login action will fire and it will get the login model containing the username and password the user entered.

It's then validated against user records to check if it's valid login if not return 401. If it is valid login we build a list of user claims that come from the data source. With that we proceed with login and create cookie with identity token. The redirect to ```/```, which will reload the SPA.

We also have ```/account/logout``` & ```/account/claims``` endpoints. These endpoints are also decorated with the authorize attribute. That means that these endpoints can only be accessed after successful login.

Everytime the backend receives an identity cookie, the middleware will automatically verify and read it and make the content available through these claims.

## Frontend
We add useUser hook with below content:
```js
import React, { useCallback, useEffect, useState } from "react";
import useGetRequest from "./useGetRequest";

export const useUser = () => {
  const { get, loadingState } = useGetRequest(
    "/account/getUserClaims?slide=false"
  );
  const [claims, setClaims] = useState();
  const [isAuthenticated, setIsAuthenticated] = useState(false);

  useEffect(() => {
    const getClaims = async () => {
      const claims = await get();
      setClaims(claims);
      claims && claims.length > 0 && setIsAuthenticated(true);
    };
    getClaims();
  }, [get]);
  const login = useCallback(
    () => window.location.replace("/account/login"),
    []
  );
  const logout = useCallback(
    () => window.location.replace("/account/logout"),
    []
  );
  const getNameClaim = useCallback(
    () => claims?.find((claim) => claim.type === "name").value,
    [claims]
  );
  return {
    claims,
    loadingState,
    isAuthenticated,
    getNameClaim,
    login,
    logout,
  };
};
```

We have to makesure login logout etc are browser redirects not react's router redirect. We shouldn't use react's router to do this because then the redirects are handled internally by React. There is also a helper function to get the name claim of the user in an easy way.

Now that we have the hook it's easy to make a component that either shows a login button when user's aren't authenticated or a greeting with the logout button if they are. Call the component **Authenticator**. It calls the **useUser** hook.

```js
import loadingStatus from "@/helpers/loadingStatus";
import { useUser } from "@/hooks/useUser";
import React from "react";
import { LoadingIndicator } from "./loadingIndicator";

export const Authenticator = () => {
  const { isAuthenticated, login, logout, getNameClaim, loadingState } =
    useUser();
  if (loadingState === loadingStatus.isLoading) {
    return <h4>Loading...</h4>;
  }
  if (isAuthenticated) {
    var username = getNameClaim();
    return (
      <div>
        Hi {username}
        <div>
          <button onClick={logout} className="mt-3 btn btn-secondary btn-sm">
            Logout
          </button>
        </div>
      </div>
    );
  } else {
    <button onClick={login} className="btn btn-primary">
      Login
    </button>;
  }
};
```

### Cookies
The browser stores all cookies set by a particular website. It can store many cookies from many different sites but only if requests are made to where the cookie came from, cookies will be send along.
No other site will receive the cookies, and that sounds very safe however, there's a problem. Let's assume I am an admin and I'm logged into globomantics.com. So the browser has my identity cookie stored. In the meantime, I see a message on social media that promises me a concert ticket to a very popular band by visiting the freetickets.com website. On this website we see a page whose html looks like below:
```html
<h1>You have won!</h1>
<form action="https://globomantics.com/approve" method="post">
    <input type="hidden" name="proposalId" value="4237" />
    <input type="submit" value="Get your free ticket" />
</form>
``` 

It submits the form to the Globomantics URL that approves a proposal and the proposal ID that has to be approved is sent along. The input that contains proposalId is hidden, so users don't see it. They just see the button. So when we click the button, proposal 4237 is approved without us knowing it. And it works because I am the conference organizer and I have the identity cookie for globomantics.com. This is called **Cross Site Request Forgery** or CSRF. Which is very common and can be done in lot of ways. This is just one example.

The attack works because cookies are also sent along across sites. Another site in this case is freetickets.com, can do all kinds of requests to globomantics.com and the browser is happy to send it all cookies it has for each time. To solve this, a cookie can now have a SameSite flag that can have any of these values.
- **strict**: Means cookie will never be sent cross-site.
- **Lax**: Introduces one exception to above rule. When a get request is done, cookies will be sent along as normal. This is to accomodate hyperlinks in the HTML of other sites. So in case of Globomantics, I could for example add a link to a certain conference in a blog post that is posted to another site and that link will send along the identity cookie as normal.
    - The reasoning behind this exception is that a get request should only get data not manipulate it. So be very sure that all your get endpoints only get data. It doesn't matter in which application. Under no circumstances should they add or update something. 
    - When our backend application sets the identity cookie, it will by default use SameSite cookies with a lax setting. But in our BFF scenario, we want to make sure that other sites can't access our endpoints at all. So we deviate from the default on the server, setting the strict SameSite cookie option on the configuration object.
    - We also override the default cookie name by a name that starts with __Host, and this is an extra protection that makes sure even subdomains can't exploit the cookie.
    - Also the default behaviour of the cookie middleware where users are not authenticated is to redirect to a login endpoint. But our frontend application doesn't expect HTML when authentication fails when the **getUserClaims** endpoint is called. 

## Deployment & Debugging
We have both backend and frontend in place now let's start them both. Do you see the problem here?
Both applications are running on localhost but on a different port, which is a different site. With samesite cookies on, the backend application will set the cookie in the browser when authentication is complete but the browser will never send it along in subsequent requests because it's a SameSite cookie. 
The browser has the cookie for the URL the backend runs on but when the frontend application does a request to the backend it will not send it because the frontend is running on a different domain. To tackle this one, we need to introduce some reverse proxying.

We are still going to start both applications using their own URLs but the browser will not access the frontend directly instead it will send all requests to the backend. The backend will check if it has an endpoint configured for the request and if so handle it.

So for example a request to ```/account/login``` will send back the login page to the browser. The request to ```/houses``` will return a list of houses. But for all requests that don't have an endpoint configured, the call will not be handled by the backend itself but instead sent over to the frontend application.

So when the browser simply accesses localhost:7180 the root of the domain, the backend doesn't have an endpoint configured for it and proxies the request to the React application that will generate a response containing all HTMLs users in JavaScript needed for the SPA to run.

<img width="828" alt="Screenshot 2025-01-03 at 5 09 15 PM" src="https://github.com/user-attachments/assets/34665960-89b3-4407-b57f-068f916cb6a5" />

The backend will simply pass the response on to the browser. So since the browser now only communicates with the backend, samesite cookies will work. So our backend is spawning frontend and sending request to it. For example in dotnet it can be configured in csproj file as below:
```csproj
<Project>
    <PropertyGroup>
        <SpaRoot>..\frontend\</SpaRoot>
        <SpaClientUrl>http://localhost:3000</SpaClientUrl>
        <SpaLaunchCommand>npm run dev</SpaLaunchCommand>
    </PropertyGroup>
</Project>
```

Another benefit is we don't have to deal with CORS in this case. React application does the API request directly to the API using relative URL. No need to put in domain here because it runs on the same one as the backend application and because of that backend doesn't have to be configured with the CORS policy.

This reverse proxy built into the backend is ideal while developing the application because we can debug everything at once, including authentication. However, in production a scenario like this might be more suitable.

<img width="1230" alt="Screenshot 2025-01-03 at 5 07 55 PM" src="https://github.com/user-attachments/assets/6bc270a0-cbbf-493c-b919-672f751fec28" />

Here a separate reverse proxy using a webserver like NGINX or IIS for example runs on app.globomantics.com and it is proxying both requests for the backend and the frontend to their corresponding URLs.
In this setup, we could deploy the static assets for the React application on the CDN, for example, that runs on a domain that is different from the one the backend application uses. Because from the browser's perspective, cookies set by the backend come from app.globomantics.com and the React application is coming from the same domain, SameSite cookie will work.

Please visit https://bit.ly/reverseprox to see an example on how this is setup.

# OpenID Connect & OAuth2
**Identity providers** provide authentication services to multiple applications. Modern ones implement the industry standard **OpenID Connect** to make that work. Expect to know all about them in combination with BFF by the end of this module.

When the application landscape in your organization has similarities with picture below, using just cookie authentication is probably not enough.

<img width="1184" alt="Screenshot 2025-01-03 at 5 23 07 PM" src="https://github.com/user-attachments/assets/ae019493-4479-4245-8ca7-585f0729192b" />

Imagine we have two frontends for the same application one on the web and one on the phone. We could implement BFF for both and do authentication in there. But users will expect they can use the same credentials and in addition to that in above diagram both frontends use separate web apis to get and manipulate data. But the API has to have a way to restrict access to it. The API needs to make sure the user was authenticated in the frontend application and it wants to know who the user is maybe because it has to filter data acoording to the role claim of the user for example.

We need some mechanism to make that work and just using cookies won't suffice because cookies are domain-restricted, and all these applications run separately.

For more centralized approach to authentication, an **identity provider** is used. An identity provider is a service application. You could see it as a special kind of API that is added to the application landscape. It facilitates a centralized way to do authentiation that works across several applications.

Authentication doesn't take place at one specific application but instead at this centralized identity provider and the user store with claims is managed by the identity provider too. When user want to change their password setup two factor authentication or do something else to manage their account, they do it at the identity provider as opposed to the application they're in.

There's typically only one identity provider for all applications in an organization and that brings another benefit to the table, when users use different web applications that all trust the same identity provider, they have to login only once in order to use all these applications and this effect is known as **Single Sign-On (SSO)**. 

This time authentication for React app is not done by the backend companion but instead by the identity provider. The backend just facilitates it and that means that login page is now not displayed by the backend, instead it will redirect to the identity provider, which will handle the login process and the datastore with all user claims we used earlier is moved to it because identity providers are the source of truth for claims.

Once a user logs into the identity provider an identity cookie is set in the browser. The identity provider will then let the backend know that the login was successful by sending it the user information, the user claims. Since it now has the user's claims, it can set its own identity cookie as normal. On subsequent requests the identity provider is not needed anymore because we can just rely on the identity cookie.

Now let's say the browser browses to another application, an aplication within the same organization. That application also redirects to the same identity provider but the browser already has an identity cookie from the identity provider, so no need to login. The identity provider skips the login process and sends the user information to the application striaghtaway, which will set the cookie for that particular application in the browser too. And this is called single sign-on.

<img width="1137" alt="Screenshot 2025-01-03 at 6 01 55 PM" src="https://github.com/user-attachments/assets/073296ea-24f0-4d71-8be3-c26bcc95becf" />

The only piece of information missing from above picture is how the identity provider communicates with our backend to send the user information.

## Identity Provider Concepts
Consider the example of Hotel, we are there for hotel reservation. The will ask for an ID. Let's say it's passport to check the claims that are on it. But there is more to it.
Everybody could fabricate a document with some claims and that's why the passport is issued by an **authority**, an organization that keeps your personal information safe. The passport is only issued by the authority with your permission. The authority in this case might be the government of the country you're living in.

The hotel knows that the passport is issued by the authority because the document says so, to prevent others from fabricating similar document or copying the document. The document also has certain characteristics. It could be a stamp or a hologram mark or a certain kind of paper. But that is not enough for the hotel to know for sure that you are who you say you are. There also has to be a trust relationship between the hotel and the authority. 
The hotel has to trust that the personal information on file at the government and thus the claims in the document, are the truth. This trust relationship is the reason the hotel doesn't accept your golf club membership card as an ID document.

There is also a trust relationship between you and the issuer of the document because you trust the issuing organization to keep your information safe.

Now we are taking this analogy to the world of applications, applications protected by an **identity provider**. In our analogy, the hotel was the party that needed authentication. In the application world, an application needs authentication, and such an application is called a client, so our backend application is a client. 

Just like the hotel had a trust relationship with the issuing organization of the passport, a client has a trust relationship with the identity provider. And just as you as the guest of the hotel, had that same trust relationship, the user of a client has it too.

The issuer of your passport keeps all your information safe. Same thing for the identity provider. The identity provider protects the personal data in the form of claims and also a second type of resource access to an API on behalf of you, the user. And ofcourse to make sure eavesdroppers have no chance, all identity provider interaction is done via HTTPS.

That's great but somehow this information has to be sent to the client. We need something like the passport was for the hotel, something that can show the client claims while at the same time client knows for sure it can trust these claims to be true and that's where tokens come in.

pic

Tokens are issued by the identity provider, but only with your permission and not before certain cryptographic characteristics are added to them. The digital version of the anti-forgery measures the passport has. The token is then sent to the client, which verifies the token using these characteristics. It makes sure the token came from the identity provider it trusts, and it checks if the token wasn't tampered with along the way.

There are two types of tokens corresponding to the two types of resources the identity provider protects. The identity token contains the personal data of the user, in other words, the claims.

It is consumed by the client as soon as it receives it. It uses the information contained to create the identity cookie and then discards it. So notice the difference between an identity token and identity cookie here. These are two separate things.

Now when it's time for the client to do a request to a protected API it has to send along the access token. Now, I am not talking about the embedded backend endpoints we've used so far, but about the APIs that are separate applications, allowing them to be accessed by potentially many applications in the application landscape.

The access token serves two purposes it acts like key that gains access to the API and it can also contain one or more claims of the user so the API can do authorization if needed. You don't need an access token per API, by the way. 