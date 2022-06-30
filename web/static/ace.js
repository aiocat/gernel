// Copyright (c) 2022 aiocat
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

var editors = [];
var currentTab = -1;

const addTab = () => {
  let editorId = editors.length;
  let clickedRemove = false;

  let tabDiv = document.createElement("div");
  tabDiv.id = `tab${editorId}`;
  tabDiv.className = "tab";
  tabDiv.innerText = `Script ${editorId + 1}`;
  tabDiv.onclick = () => {
    if (clickedRemove) return;
    currentTab = editorId;
    setEditor();
  };

  let removeTab = document.createElement("button");
  removeTab.className = "remove-tab";
  removeTab.innerText = "Remove";
  removeTab.onclick = () => {
    if (editorId === 0) return;

    clickedRemove = true;
    currentTab = 0;
    tabDiv.remove();
    document.getElementById(`editor${editorId}`).remove();
    editors[editorId] = null;
    setEditor();
  };

  let aceEditor = document.createElement("div");
  aceEditor.className = "editor";
  aceEditor.id = `editor${editorId}`;
  document.getElementById("editor-zone").appendChild(aceEditor);
  editors.push(ace.edit(aceEditor.id));

  tabDiv.appendChild(removeTab);
  document.getElementById("tabs").appendChild(tabDiv);

  currentTab = editorId;
  setEditor();
};

const setEditor = () => {
  editors[currentTab].setTheme("ace/theme/dracula");
  editors[currentTab].session.setMode("ace/mode/lua");
  editors[currentTab].setShowPrintMargin(false);

  for (let index = 0; index < editors.length; index++) {
    if (editors[index] === null) continue;

    document.getElementById(`tab${index}`).style.borderLeft = "none";
    document.getElementById(`editor${index}`).style.display = "none";

    if (index === currentTab) {
      document.getElementById(`tab${index}`).style.borderLeft =
        "6px solid #f2cdcd";
      document.getElementById(`editor${index}`).style.display = "block";
    }
  }
};

const saveFile = () => {
  let luaBlob = new Blob([editors[currentTab].getValue()], {
    type: "text/lua",
  });
  let blank = document.createElement("a");
  blank.download = "unknown.lua";
  blank.href = URL.createObjectURL(luaBlob);

  document.body.appendChild(blank);
  blank.dispatchEvent(new MouseEvent("click"));
  blank.remove();
  URL.revokeObjectURL(luaBlob);
};

addTab();
document.getElementById("add-tab").onclick = addTab;
