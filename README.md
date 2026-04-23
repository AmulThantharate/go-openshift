### 1. Initialize Git Repository

```bash
git init
```

### 2. Add Files to Git Repository

```bash
git add .
```

### 3. Commit Files to Git Repository

```bash
git commit -m "first commit"
```

### 4. Create GitHub Repository

```bash
gh create
```

### 5. Connect Local Repository to GitHub Repository

```bash
git remote add origin <github_repository_url>
```

### 6. Push Files to GitHub Repository

```bash
git push -u origin main
```

### 7. Install oc CLI

```bash
sudo dnf install openshift-clients
```

### 8. Login to OpenShift Cluster

```bash
    oc login --token= --server=
```
