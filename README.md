# swifty.demo - Functions and demo app for Swifty

Swifty is a serverless platform that allows you to develop and run backend for your application as fast and easy as possible without any infrastructure management. With Swifty you do not need to think about scalability or availability of your application - we do it for you. We also include Maria and Mongo databases, Object Storage and Authentication-as-a-Service for your comfort. You only need to code functions for your backend.

How to use Swifty for your web or mobile application?

In this example we will use Swifty as a backend for <a href="https://demo.swifty.cloud">swifty.demo</a> portal. Swifty.demo is a simple application that allows you to manage user profiles. We will use Go and Python to build functions while swifty.demo page was created with vue.js. You can find all functions and vue.js code here <a href="https://github.com/swiftycloud/swifty.demo">https://github.com/swiftycloud/swifty.demo</a>

You can find more functions here https://github.com/swiftycloud/swifty.demo/tree/master/functions

# Swifty Deployment

First, you need to get access to Swifty. Swifty is available from our own <a href="https://dashboard.swifty.cloud" target="_blank" rel="noopener">Swifty.Cloud</a>, on-premises or you can deploy it to your favourite public cloud. We have created an Ansible script to allow you deploy Swifty to Amazon AWS. <a href="https://swifty.cloud/2018/06/20/how-to-deploy-swifty-to-aws/">Here is the guide how to do it</a>. This version of Swifty called Swifty.Lite and allows you to use all our features but with some limitations.

After deployment you need to create a user in Swifty using swifty.dashboard.

# Authorisation and Profile Management

For typical application that stores our data we, obviously, need an ability to create and manage user accounts. Swifty has integrated Authentication-as-a-Service solution that allows you to perform basic auth actions: SignUp, SignIn and Logout (leave).

First, we need to create authentication backend. Please go to Authentication Services in swifty.dashboard and Create Auth Database. You need to name it "demo_auth".

Authentication service consists of three parts:
<ol>
 	<li>MongoDB that stores all data.</li>
 	<li>Function that serves authentication calls. Default function provides basic features you need to manage users and you can use it as is without any changes. But if you need an additional features or integration with 3d-party services for your authentication service you can easily update the provided function as you wish, for example, to add integration with mail services. Function also equipped with two env variables: SWIFTY_AUTHJWT_MONGO containing the name of mongo mware and SWIFTY_AUTHJWT_MWARE containing the name of auth mware.</li>
 	<li>Auth mware that allows you to enable authenticated calls for functions and manages JWT tokens. Currently only the "authjwt" one is supported. In swifty library there's a helper that makes use of this mware by generating (and signing) JWT tokens with the given claims.</li>
</ol>

Function includes URL Trigger that should be called by application for user management actions and should have access to mongo and auth mwares. To get URL trigger link please go to Functions, find auth function (marked as Authentication), go to Triggers and copy REST API URL. Please go to <a href="https://demo.swifty.cloud">swifty.demo</a> and past your Auth REST API URL to Auth URL field.

The expected arguments for function are "action" which can be "signup”, "signin" or “leave”, "userid" and "password”. Password is kept in a secure way, it's salty-hashed with 256-bit salt and bcrypt slow hasher.

For example, to add new user you need to place the following call using POST method:
<blockquote>POST https://api.dtln.ru:8686/call/012cb5e1ee010032dce5d8f5f9d8c90239d32956bae998892d9db7a2b3c22bf7?action=signup&amp;userid=user@yourmail.com&amp;password=xxxxxxxx</blockquote>

to sign in:
<blockquote>POST https://api.dtln.ru:8686/call/012cb5e1ee010032dce5d8f5f9d8c90239d32956bae998892d9db7a2b3c22bf7?action=signin&amp;userid=user@yourmail.com&amp;password=xxxxxxxx</blockquote>

to logout:
<blockquote>POST https://api.dtln.ru:8686/call/012cb5e1ee010032dce5d8f5f9d8c90239d32956bae998892d9db7a2b3c22bf7?action=leave&amp;userid=user@yourmail.com</blockquote>

If signin is successful you will get a JSON response with JWT "token”. Using this token you will be able to access functions with authentication enabled. The JWT token is to be provided via the <a href="https://swagger.io/docs/specification/authentication/bearer-authentication/">Authorization Bearer scheme</a>. If auth checks do not pass (no header, not bearer scheme or bad signature) the function is not called, user gets 401 code back. If passed, the function receives the token claims as the _SWY_JWT_CLAIMS_ argument.

Our swifty.demo app do the following. When you SignUp and enter user data the page:
<ol>
 	<li>SignUp with new user using provided AUTHaaS.</li>
 	<li>Get JWT token and create new record in user database with provided user data.</li>
 	<li>SignIn with new user.</li>
 	<li>Get user profile from the database.</li>
</ol>

To use swifty.demo you need also to create a function to manage user profiles and database to store them. First we will create profile management function. Please go to Swifty and create it from template: (create template for profile management).

<ol>
 	<li>Go to Functions -&gt; New Function -&gt; Name it “demo_profile_management” -&gt; Select Go language -&gt; Select template: Profile Management</li>
 	<li>Go to Functions -&gt; demo_profile_management -&gt; Triggers -&gt; Add Trigger -&gt; REST API (URL)</li>
 	<li>Copy just created URL and past it to the Profile Function URL field of swifty.demo.</li>
</ol>

Second we need to create MongoDB and assign it to profile management function.
<ol>
 	<li>Go to MongoDB -&gt; New Database -&gt; Name it “demo_profiles"</li>
 	<li>Go to Functions -&gt; demo_profile_management -&gt; Middleware -&gt; Attach Middleware -&gt; Select MongoDB and “demo_profiles” instance.</li>
</ol>

# Profile Picture Management

Swifty.demo also allows you to upload profile picture. Swifty has integrated Object Storage to store files, pictures and any other objects. It has Amazon S3 compatible API, so you can, actually, manage and use it with any S3 management tool. We will use it to store user profile pictures.

To create object storage for our project:
<ol>
 	<li>Go to Object Storage -&gt; Create Bucket -&gt; Name it “demo_images".</li>
 	<li>Go to Functions -&gt; demo_profile_management -&gt; Middleware -&gt; Attach Middleware -&gt; Select Object Storage and “demo_images” instance.</li>
</ol>

To upload and download profile picture you can use standard S3 SDK for your favourite programming language. But we created special function that allows you to upload and get picture from Object Storage easily.
<ol>
 	<li>Go to Functions -&gt; New Function -&gt; Name it “demo_avatar_management” -&gt; Select Python language -&gt; Select template: Avatar Management</li>
 	<li>Go to Functions -&gt; demo_avatar_management -&gt; Triggers -&gt; Add Trigger -&gt; REST API (URL)</li>
 	<li>Copy just created URL and past it to the Avatar Function URL field of swifty.demo.</li>
</ol>

Great! Now you have all you need to create your first user. Please go to <a href="https://demo.swifty.cloud">swifty.demo</a> and SignUp with new user credentials. Now you are logged into the portal and you see your user profile. Please change City field, upload profile picture and click Update button.

If you go to Functions screen you will see that your auth function was executed at least one time to create new user account. Your demo_profile_management function also was executed to create user profile. Please also go to your Object Storage and you will find your picture in demo_images bucket.

# SignIn and LogOut

Swifty.demo portal, of course, also allows you to LogOut and SignIn back. You can try it if you want.

# Next Steps

Now you have all basic information to start using Swifty services. You can find all functions and swifty.demo code in this repo.
More test functions are available here https://github.com/swiftycloud/swifty.demo/tree/master/functions

If you have any question please <a href="mailto: info@swifty.cloud">ask us</a>.
