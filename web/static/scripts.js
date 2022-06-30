// Copyright (c) 2022 aiocat
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

const loadScripts = async () => {
  document.getElementById("scripts").innerHTML = "";

  let content = await (await listScripts()).text();
  for (const luaScript of content.split("\n")) {
    if (!luaScript) continue;

    let button = document.createElement("button");
    button.innerText = luaScript;
    button.onclick = async () => {
      let luaContent = await (await getScript(luaScript)).text();
      editors[currentTab].setValue(luaContent);
    };

    document.getElementById("scripts").appendChild(button);
  }
};

loadScripts();
