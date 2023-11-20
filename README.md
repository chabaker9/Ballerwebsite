# History of the Internet

Welcome to 'The History of the Internet,' a digital journey that explores the remarkable evolution of the internet since its inception in 1969. This interactive experience explores the internet's origins with ARPANET, its pivotal role in revolutionizing global communication and commerce, and its ongoing impact on society. From the pioneering days of basic connectivity to the current era of social media, e-commerce, and IoT, our website offers a detailed look at how the internet has transformed our world, connecting billions of people and reshaping every aspect of modern life.

This website is the product of my 11th grade "Modern Web Development" class which was a great exercise allowing me to put to use all the concepts I learned in class.

## Components

### Content

Standard HTML with CSS with back end templating.

### License file

A [LICENSE](src/LICENSE) file file in open-source code clearly outlines the terms under which others can use, modify, and distribute your work, providing legal clarity and protecting your rights as the creator. It's essential to ensure your work is used in ways you approve of and to encourage its wider adoption and contribution.

The Apache License 2.0 is a permissive, flexible option that allows extensive freedom in code usage, including in proprietary projects, while ensuring proper attribution to the original creator. Additionally, it explicitly grants patent rights, reducing the risk of patent litigation for both contributors and users.

### Web server

Even though the website contains mostly static content, I decided to use a [Go based web server](src/main.go) so I could experiment with backend [templating](https://github.com/chabaker9/Ballerwebsite/blob/main/src/templates/info.html#L11-L26)

### Deployment

The website is deployed to Google Cloud Platform's (GCP) with two steps

1. Build container image from [Dockerfile](Dockerfile) -  `gcloud builds submit --tag $image`
1. Deploy container image to Cloud Run - `gcloud run services delete $appName --platform managed --region us-central1`

### Visit the finished product

[History of the Internet](https://ballerwebsite-sdhhrs6v5a-uc.a.run.app)
