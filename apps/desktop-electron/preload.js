const { contextBridge } = require("electron");

contextBridge.exposeInMainWorld("aiNailsShell", {
  platform: process.platform,
  shell: "electron",
  packaged: process.env.NODE_ENV === "production",
});
