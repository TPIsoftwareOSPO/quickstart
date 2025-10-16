# quickstart

**quickstart** is a **lightweight command-line orchestration tool** built with Go and Cobra.  
It helps developers **automate, organize, and monitor multi-command workflows** through a clear, declarative YAML configuration.

Instead of juggling multiple terminals or maintaining long shell scripts, **quickstart** lets you:  
- Define all your tasks in one YAML file.  
- Control dependencies between processes.  
- Run everything with **a single command**.  

Perfect for setting up **local dev environments**, running **integration tests**, or managing **automation pipelines** — all in a fast, reliable, and reproducible way.

---

## 🚀 Why quickstart?

- ⚡ **Simplify your setup** – Stop memorizing long command chains. Just define them in YAML.  
- ✅ **Reliable orchestration** – Tasks run only when their dependencies are ready.  
- 🔄 **Flexible execution** – Run concurrently or sequentially as your workflow requires.  
- 🧑‍💻 **Developer-friendly** – Ideal for repeatable setups in dev, test, or CI/CD environments.  
- 🔍 **Transparent & observable** – Easy to debug, verify, and reproduce results.  

---

## 💡 Typical Use Cases

- Spin up **local environments** with multiple services (databases, APIs, message queues).  
- Run **integration tests** that depend on multiple running services.  
- Automate **CI/CD workflows** with predictable, declarative orchestration.  
- Replace fragile **bash scripts or Makefiles** with structured YAML.  


---

## 🧩 Key Features

* 🧾 **Declarative YAML config** – Define tasks, dependencies, and health checks cleanly.
* ⚙️ **Process orchestration** – Manage concurrent or sequential command execution.
* ❤️ **Built-in health checks** – Validate task readiness before proceeding.
* 🌐 **HTTP & JSON path validation** – Verify responses and extract data easily.
* 💻 **Command-based checks** – Use shell commands to test custom health states.
* 🧠 **Dependency control** – Define startup order via `depends_on`.
* 💡 **Flexible execution** – Customize `base_dir`, `args`, and `envs` per task.

---

## 🔍 Comparison to Alternatives

| Tool               | Key Difference                                              |
| ------------------ | ----------------------------------------------------------- |
| **Shell Scripts**  | YAML offers structure, dependency logic, and health checks. |
| **Makefiles**      | Designed for orchestration, not just builds.                |
| **Docker Compose** | Works with any executables — not just containers.           |

---

### 🏁 Summary

**quickstart** brings structure, reliability, and observability to your automation workflows —
helping you focus more on building and less on managing scripts.





# $\color{Red}\Huge{\textsf{Important Note!}}$

Because we have $\color{Red}\large{\textbf{NOT}}$ purchased Apple and Microsoft developer accounts,
you'll currently need to manually bypass system security blocks when running the application.


## Installation

### Using Package Manager

#### Homebrew (macos/linux)

- [homebrew installation](https://brew.sh/)

```shell
brew install --cask TPIsoftwareOSPO/tap/quickstart

# macos (as "Important Note!" mention above)
xattr -dr com.apple.quarantine $(which quickstart)
```

#### Scoop (windows)

- [scoop installation](https://scoop.sh/)

```shell
# add scoop tap
scoop bucket add TPIsoftwareOSPO https://github.com/TPIsoftwareOSPO/scoop-bucket.git
# install quickstart
scoop install quickstart
```

### Portable Executable

Download the compressed file that matches your system from [release](https://github.com/TPIsoftwareOSPO/quickstart/releases).
The filename prefix will be `quickstart-portable_`.

To run the portable executable, it must be in the same directory with the `quickstart.yaml` file.

### Basics

`quickstart.yaml`

```yaml
tasks:
  - name: echo
    executable: echo
    args:
      - hello
      - world
```

To execute tasks defined in your configuration file:

```bash
quickstart up
```

output:
```shell
quickstart|Using config file: ...(ellipsis)/quickstart.yaml
echo|hello world
echo|Completed
```

### Using Docker

```shell
docker run --rm -v ./quickstart.yaml:/app/quickstart.yaml tpisoftwareopensource/quickstart up
```

### Command Line Interface (CLI)

`quickstart` provides a straightforward command-line interface.

**Usage:**

```bash
quickstart [command]
```

**Available Commands:**

| Command    | Descriptions                                                |
|:-----------|:------------------------------------------------------------|
 | check      | Confirm the correctness of the YAML content format.         |
 | completion | Generate the autocompletion script for the specified shell. |
 | down       | Kill previous tasks processes.                              |
 | help       | Help about any command.                                     |
 | up         | Execute tasks according to the YAML configuration file.     |
 | version    | Show version number and build details of quickstart.             |
| init       | Generate minimal quickstart.yaml file                            |


### Configuration File Reference

The `tasks` key at the root of your YAML file contains a list of individual task definitions. Each task can have the following properties:


| key                                     | type               | description                                                                                                                                                      |
|:----------------------------------------|:-------------------|:-----------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `name`                                  | string, required   | A unique identifier for the task.                                                                                                                                |
| `base_dir`                              | string             | The working directory for the command. cmd.Dir will be set to this path. If not specified, the current working directory of quickstart will be used.                  |
| `executable`                            | string, required   | The path to the executable command                                                                                                                               |e.g., node, java, ./my-app.|
| `args`                                  | []string           | A list of arguments to pass to the executable.                                                                                                                   |
| `envs`                                  | []string           | A list of environment variables to set for the command                                                                                                           |e.g., KEY=VALUE. These are merged with the parent process's environment variables.|
| `depends_on`                            | []string           | A list of task names that this task depends on. This task will only start after all its dependencies have successfully passed their health checks.               |
| `healthcheck`                           | object             | Defines how quickstart determines if a task is healthy.                                                                                                               |
| `healthcheck.http`                      | object             | Configures an HTTP GET health check.                                                                                                                             |
| `healthcheck.http.url`                  | string, required   | The URL to send the HTTP GET request to.                                                                                                                         |
| `healthcheck.http.expect`               | object, optional   | Defines expected responses.If not set, a 2xx HTTP status code indicates health.                                                                                  |
| `healthcheck.http.expect.json`          | object             | Expects a JSON response.                                                                                                                                         |
| `healthcheck.http.expect.json.jsonpath` | string, required   | A JSONPath expression to extract a value from the response.                                                                                                      |
| `healthcheck.http.expect.json.value`    | string, required   | The expected value                                                                                                                                               |as a string to match against the extracted JSONPath value.|
| `healthcheck.command`                   | object             | Configures a command-based health check.                                                                                                                         |
| `healthcheck.command.scripts`           | []string, required | A list where the first element is the command, and subsequent elements are its arguments. The command is considered healthy if it exits with a zero status code. |
| `healthcheck.frequency`                 | object             | Controls the timing of health checks.                                                                                                                            |
| `healthcheck.frequency.interval`        | duration string    | The time between consecutive health check attempts                                                                                                               |e.g., 5s, 1m.|
| `healthcheck.frequency.timeout`         | duration string    | The maximum time allowed for a single health check attempt                                                                                                       |e.g., 10s.|
| `healthcheck.frequency.retries`         | int                | The maximum number of consecutive failed health checks before the task is considered unhealthy.                                                                  |
| `healthcheck.frequency.delay`           | duration string    | The initial delay before the first health check attempt is made after a task starts                                                                              |e.g., 5s.|

### Logging

The application's startup logs will be located in the `logs/` directory. 

Each log file will be named in the format `{task:name}-{date}.log`

### Examples

#### Example 1: Java Application Portable Launch Package

Here's an example of how you might configure `quickstart` to start .jar file using java command

**Prerequisite**
1. prepare your `myapp.jar` file.
2. (Optional) prepare portable jre. (e.g [Azul Zulu](https://www.azul.com/downloads/#zulu))
3. create a new directory, put in your `myapp.jar` and jre directory. The file structure as below:

```
new-dir/
├─ jre/
    ├─ bin/
        ├─ java(.exe)
├─ myapp.jar
├─ quickstart.yaml
├─ (Optional) quickstart-portable(.exe)
```

**Example of quickstart.yaml**

- With Portable Jre Java
    ```yaml
    tasks:
      - name: myapp
        base_dir: .
        executable: jre/bin/java(.exe) # portable jre executable
        args:
          - -jar
          - myapp.jar
    ```

- Use System Jre Java
    ```yaml
    tasks:
      - name: myapp
        base_dir: .
        executable: java # call system java command
        args:
          - -jar
          - myapp.jar
    ```

Double-click the quickstart portable executable or run `quickstart up` in your current directory's terminal.

#### Example 2: Portable Elasticsearch and Kibana Set

Here's an example of how you might configure `quickstart` to start Elasticsearch and Kibana, 
with proper health checks and dependencies:

**Prerequisite**

1. download [elasticsearch](https://www.elastic.co/downloads/elasticsearch)
2. download [kibana](https://www.elastic.co/downloads/kibana)
3. Unzip the downloaded files and place them in the same directory. The file structure as below:

files structure:
```
new-dir/
├─ elasticsearch/
├─ kibana/
├─ quickstart.yaml
├─ kibana.yml
```

**Example of kibana.yaml**

```yaml
csp.strict: false
server.ssl.enabled: false
telemetry:
  enabled: false
xpack:
  encryptedSavedObjects:
    encryptionKey: "01234567890123456789012345678901" # only for demo
```

**Example of quickstart.yaml**

```yaml
tasks:
  - name: elasticsearch
    base_dir: ./elasticsearch
    executable: bin/elasticsearch
    args:
      - -E
      - xpack.security.enabled=false
      - -E
      - xpack.security.http.ssl.enabled=false
      - -E
      - xpack.security.transport.ssl.enabled=false
      - -E
      - xpack.monitoring.collection.enabled=true
    healthcheck:
      frequency:
        interval: 5s
        timeout: 10s
        retries: 5
        delay: 5s
      http:
        url: http://localhost:9200
  - name: kibana
    base_dir: ./kibana
    executable: bin/kibana
    args:
      - -c
      - kibana.yml
    healthcheck:
      frequency:
        interval: 5s
        timeout: 10s
        retries: 5
        delay: 5s
      http:
        url: http://localhost:5601/api/status
        expect:
          json:
            value: available
            jsonpath: "$.status.overall.level"
    depends_on:
      - elasticsearch
  - name: curl1
    executable: curl
    args:
      - -v
      - http://localhost:9200
    depends_on:
      - elasticsearch
  - name: curl2
    executable: curl
    args:
      - -v
      - http://localhost:5601/api/status
    depends_on:
      - kibana
```

Double-click the quickstart portable executable or run `quickstart up` in your current directory's terminal.

# Implementations

- [how-to-make-digirunner-opensource-launch-standalone](https://github.com/vulcanshen-tpi/how-to-make-digirunner-opensource-launch-standalone)
