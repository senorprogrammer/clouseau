<pre>
 ██████╗██╗      ██████╗ ██╗   ██╗███████╗███████╗ █████╗ ██╗   ██╗
██╔════╝██║     ██╔═══██╗██║   ██║██╔════╝██╔════╝██╔══██╗██║   ██║
██║     ██║     ██║   ██║██║   ██║███████╗█████╗  ███████║██║   ██║
██║     ██║     ██║   ██║██║   ██║╚════██║██╔══╝  ██╔══██║██║   ██║
╚██████╗███████╗╚██████╔╝╚██████╔╝███████║███████╗██║  ██║╚██████╔╝
 ╚═════╝╚══════╝ ╚═════╝  ╚═════╝ ╚══════╝╚══════╝╚═╝  ╚═╝ ╚═════╝ 
                                                                   
</pre>

A configuration inspector/checker for Ruby on Rails. 

Clouseau will analyze your Rails app and provide information about the following configuration uses.

<img src="./images/clouseau_ui.png" alt="clouseau ui example" />

## Config gem
https://github.com/railsconfig/config

* Analyze config's `settings.yml` files to show you all the configured variables and values
* Indicate which values are missing from the base configuration file (aka `settings.yml`)
* Indicate which production values are hard-coded the same as values for other environments
* Indicate which values are derived from other configuration files, and what those derived values will be
* Indicate where in your Rails app all the Config config values are used

## Figaro gem
https://github.com/laserlemon/figaro

* Indicate where in your Rails app all the Figaro config values are used

## Environment Variables

* Indicate where in your Rails app all the environment variables are used