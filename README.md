# Application Generator
Generator to create your personal application.
Currently is only one style available but in future it will increase.

To see how it works, please have a look into the examples.

Basically you have to fill a `genmodel.application` struct and
call the method: `generator.Build(application, "output-path")`

Inside your profile you can set every necessary field for the application,
inside the jobposition you can choose some config things:
- color
- language
- style

and you can specify your profile content depending to the job offer:
- tech skills
- soft skills
- motivation text

## Supported Link Icons (will increase in the future)
- https://github.com
- https://www.linkedin.com
- https://www.xing.com

## Supported Languages
- English
- German