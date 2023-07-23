### Kapitan
---
##### Kapitan Definition
```
Kapitan is a tool that allows for generically managing configurations for various applications in a DRY manner. Kapitan 
can configure tools such as Kubernetes, Terraform, or Ansible but is not biased towards any application.
```

### Kapitan Inventory
---
##### Inventory
```
Inventory is a hierarchical YAML structure used to capture configurations that can be passed on to templating engines.
There will typically be a folder called "inventory" that holds classes and targets. 
```

##### Target
```
Targets are YAML files that tell Kapitan what a user would like to do with Kapitan. Typically, targets can be separately 
defined for each environment a user would like to deploy to. These files are located under the "inventory/targets" 
directory.
```

##### Class
```
Classes are reusable common configurations defined through YAML files that eliminate duplication. Targets can list 
classes in its manifest, causing their contents to be merged together. These files are located under the 
"inventory/classes" directory.
```

### Kapitan Input Types
---
##### Input Types
```
Input types are templating engines that will render the inventory upon compilation. Examples of template engines include 
Generators, Kadet, Jsonnet, and Jinja2.
```

##### Compile Phases
```
Step 1) Kapitan uses "reclass" to render a final version of the inventory. Reclass is an open source tool for merging 
        data sources recursively.
Step 2) Kapitan fetches external dependencies per parameters.kapitan.dependencies.
Step 3) Kapitan compiles the input types for each target per parameters.kapitan.compile.
Step 4) Kapitan reveals the secrets directly in the compiled output per parameters.kapitan.secrets.
Step 5) Kapitan copies output files to the /compiled directory.
Step 6) Kapitan validates the compiled output per parameters.kapitan.validate.
```