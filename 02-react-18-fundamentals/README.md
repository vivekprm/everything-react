# What Problem Does React Solve?

- Static UIs are easily manageable.
- Dynamic Changes are hard
- Lots of code needed that's difficult to maintain.

React helps, it brings structure with components.

# Core React Features

- Offers structure by letting us compose the UI with separate reusable components. Components are the building blocks of a React application and can have a changeable state.
- The UI of a component is declared in JavaScript and that's unconventional. But we'll see that this has a huge advantage. It gives React the ability to make it easier to dynamically change the UI of components when the state changes. The declared UI will be rendered to a browser or other output automatically by React.
- React is very efficient because when the UI updates, it will not just bluntly replace the whole rendered component. Instead, a smart mechanism called **reconciliation** will figure out what actually changed and it will only update that.

# Components & JSX

- Components are JavaScript functions that return JSX, a syntax that seems like HTML.
- JSX looks like HTML but it is not, it's just an alternative way to write JavaScript.
- JSX stands for JavaScript Extension. It extends the JavaScript syntax to make it easier for us developers to write UIs in an intuitive hierachical way.
- Browser doesn't understand JSX. JSX has to be transformed by a tool to JavaScript. The typical tool for that is called **Babel**. We can go to Babel website and try out how it converts JSX to JavaScript.

## JSX to Markup

- So from JSX, we can get to **createElements** statements using Babel tool.

```
<h1>This is a banner</h1> ----> React.createElement("h1", null, "This is a banner")
```

- What happens next and how does the browser know what to render. Each time, browser has to render a component a JavaScript library called React DOM uses to create element code to generate the actual HTML elements for the browser.

## Separation of Concerns

- React separates its core library from the library that is responsible for rendering components. The latter is called react-dom for browser applications but react can also be used for mobile applications using React Native.
- In case of react native the core library is still used but react-dom will be different library that can render mobile controls.

Let's look at another component:

```js
const Greeting = () => (
  <div>
    <Banner />
    <h2 className="highlight">Greetings!</h2>
  </div>
);
```

In this case, JSX returned has a div that contains the Banner component from the previous slide, it is rendered in this component using the same HTML like element syntax we used with the div and the h2. That **div and h2, in fact, are also components**. They're just built-in components that correspond to DOM elements. Since we are targetting browser here the components used will result in the rendering of HTML elements with the same name. So div for example, will result in the rendering of a div in the browser.

Now let's say you are not using a browser as a rendering target like we do in this course, but the mobile device for instance, in that case you won't use the same built-in components but the ones that corresponds to controls that are available on a mobile device.

The casing of the names of DOM related components built into React is camelcase.
Custom components created by us are Pascal Cased where the first letter is capital letter.

# Application Structure

Create project using next.js

```sh
npx create-next-app globomantics
```

# Modules

What are modules? Well they are not React specific, they are part of JavaScript.

Lets say we have `module.js` file with a function called doSomething.

```js
const doSomething = () => {
    ...
}
export {doSomething};
```

To make this file a module, we need to export something from it using the export keyword. In this case we are exporting the function. Exporting it makes it possible for other modules to import and use it. Should the file contain other things that are not exported, these things are just not available to other modules.

You can also export multiple members of the module by just separating them with commas. Members of the module can be classes, functions, object basically anything that can exist in a variable. Any module that wants to import things from this module does so using the import keyword followed by what it wants to import between curly braces.

anotherModule.js

```js
import { doSomething } from "./module";
doSomething();
```

It's also possible for a module to export one default member using the default keyword.

module.js

```js
const doSomething = () => {
    ...
}
export default doSomething;
```

Importing side then doesn't have to use the curly brace syntax and also it can use any name.

anotherModule.js

```js
import do from "./module";
do();
```

It's also possible to have combination of default export and normal export.

## Reasons to Use Modules

First of all they bring structure in the code base. Just like components are the building blocks of the UI, modules are the building blocks in the code.

- They make easy to reuse code across the application without using global variables.
- They bring Encapsulation. Everything that is not exported from a module remains private for the module.
- They are needed for bundling. The **bundler** tool will try to create one big JavaScript file out of all these seperate little files. Without modules, it would have no idea about the order of the individual files.

# Running the Finished Project

https://github.com/RolandGuijt/ps-react-fundamentals

# Adding New Components

Lets remove everything that we had generated for next.js project and add a directory called components. Add banner and app components.

banner.js

```js
const Banner = () => {
  return (
    <header>
      <div>
        <img src="./GloboLogo.png" alt="logo" />
      </div>
      <div>Providing houses all over the world</div>
    </header>
  );
};
export default Banner;
```

app.js

```js
import Banner from "./banner";

const App = () => {
  return <Banner />;
};
export default App;
```

Now add app component in index.js

```js
import App from "../components/app";

const Index = () => <App />;
export default Index;
```

# Detecting Problems & Debugging

Part of the `Next.js` tooling is **ESLint**, which is able to detect problems in the code, and it also helps with code styling.
`Next.js` comes with a ruleset for React preinstalled and you can run the tool that check these rules by typing `npm run lint` in the terminal.

When we debug we see breakpoint is hit twice and tht's because React strict mode is on. This will try to anticipate possible problems with the application and for that it executes function components twice. It won't happen in production build but if you don't like it while developing, you can turn it off in the `next.js` configuration file.

How come we see JSX code while debugging in browser? Afterall we learned that browser doesn't understand JSX.
That's because of the tooling. It creates source map files during the dev build that link the generated Babel code to the original source.

When debugging in chrome it's worthwile to install a plugin called "React Developer Tools" This will get us two extra tabs:

- Profiler
- Components

# Styling Components

There are couple of ways to style components.

- Next.js specific
  - In \_app.js CSS files can be loaded that can be applied throughout the application.
- We deleted \_app.js because we are going to load the CSS in the HTML root document.
- It's also possible to have a CSS file with styles that are isolated for just one Component.
- Use **style** attribute on a component.

Most web applications use a standard CSS file to get started style-wise.

In the demo application we are using bootstrap mainly to position the components. A typical way to load CSS file is in the HTML root document, the file with the Html & the Head tags.

But in `next.js` that's nowhere to be found in the project template. Should you use create-react-app, you will find it. It's simply a file called `index.html`. But in `next.js` it's provided behind the scenes. It's possible to access it though. Just create an \_document.js

This is the default content for it:

```js
import { Html, Head, Main, NextScript } from "next/document";

export default function Document() {
  return (
    <Html lang="en">
      <Head />
      <body>
        <Main />
        <NextScript />
      </body>
    </Html>
  );
}
```

Now we can add a link pointing to bootstrap as below:

```js
import { Html, Head, Main, NextScript } from "next/document";
import Link from "next/link";

export default function Document() {
  return (
    <Html lang="en">
      <Head>
        <link
          href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.3/dist/css/bootstrap.min.css"
          rel="stylesheet"
          integrity="sha384-QWTKZyjpPEjISv5WaRU9OFeRpok6YctnYmDr5pNlyT2bRjXh0JMhjY6hW+ALEwIH"
          crossOrigin="anonymous"
        />
      </Head>
      <body>
        <div className="container">
          <Main />
        </div>
        <NextScript />
      </body>
    </Html>
  );
}
```

## Bootstrap Grid System

The screen is divided in 12 equally wide invisible columns. When the CSS class **row** is applied to the component a new row will be created in the grid for the children of the component. In the banner component we can assing css className row to the header component.

Two divs are are the children. They can now claim a certain width of the row using a **col** class. `col-5` means 5 columns wide. so first div accupies 5 columns and second div 7 columns.

To leave some space between the row and what comes next, we can apply a second css class to header `mb-4` means margin bottom 4. So below is the complete banner.js

```js
const Banner = () => {
  return (
    <header className="row mb-4">
      <div className="col-5">
        <img src="./GloboLogo.png" alt="logo" className="logo" />
      </div>
      <div className="col-7 mt-5">Providing houses all over the world</div>
    </header>
  );
};
export default Banner;
```

## CSS Modules

- CSS modules make it possible to import a css file in a component.
- That CSS will be only available to that component as opposed to a CSS file like bootstrap that we loaded in the HTML root document. That is available for whole application for all the components.
- CSS modules are not a react feature. The functionality is provided by **webpack**, the tool that bundles CSS files for Next.js and create react app as well as others.
- With CSS modules, we can have a nicely separated CSS per component. That means when styling for a component has to be changed, there is no need to open a giant CSS file to find the classes used.
- There is no need to worry that we unintentionally change the styling of another component that uses the same class we're changing.
- But we learned all the CSS files are bundled by the tooling, so how does this work?
  - Tooling also takes care of giving all class names a unique name, and it of course will automatically apply the new names in the components.

To give the Banner component it's own css create a file called `banner.module.css`

```css
.logo {
  height: 150px;
  cursor: pointer;
}
```

Now to use this inside banner component.

```js
import styles from "./banner.module.css";

const Banner = () => {
  return (
    <header className="row mb-4">
      <div className="col-5">
        <img src="./GloboLogo.png" alt="logo" className={styles.logo} />
      </div>
      <div className="col-7 mt-5">Providing houses all over the world</div>
    </header>
  );
};
export default Banner;
```

## Style Attribute

- Just like HTML elements, components also have a style attribute.
- Difference is that in React, styles take an object that contains CSS style properties. So an expression is need to use it.
- Object can be an inline object or an existing object.

To style the banner text, we can set the style attribute on the surrounding div.

```js
import styles from "./banner.module.css";

const subtitleStyle = {
  fontStyle: "italic",
  fontSize: "x-large",
  color: "coral",
};
const Banner = () => {
  return (
    <header className="row mb-4">
      <div className="col-5">
        <img src="./GloboLogo.png" alt="logo" className={styles.logo} />
      </div>
      <div className="col-7 mt-5" style={subtitleStyle}>
        Providing houses all over the world
      </div>
    </header>
  );
};
export default Banner;
```

Use of the style attribute is however discouraged, for the same reason style attribute in HTML is discouraged.

CSS should be in separate CSS files to keep it apart from the components.
Doing it like that is also better for performance.
