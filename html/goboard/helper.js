// Configuration
const DECRYPTION_KEY = "KOK_EVM_ENCRYPTED_8888";
const API_BASE_URL = "https://m1.i247.com/kokka"; // Update this to your server URL
// const API_BASE_URL = "http://localhost:8080"; // Update this to your server URL

// Contract addresses
const CONTRACTS = {
  vndx: {
    name: "VNDX",
    address: "0x329aaF4e8d9883c6F8610D48172DE9c6C0917ecD",
    icon: "🇻🇳",
  },
  sgpx: {
    name: "SGPX",
    address: "0x6245000F860feba4619622FAF8c1eB7968cc91D3",
    icon: "🇸🇬",
  },
  yenx: {
    name: "YENX",
    address: "0xbae0597019221Fd8DB7069725F5b93B047D85a89",
    icon: "🇯🇵",
  },
};

console.log("Helper JS loaded", { DECRYPTION_KEY, API_BASE_URL });

// Encrypt private key using CryptoJS (matches server's DecryptCrypto function)
function encryptPrivateKey(privateKey) {
  try {
    const encrypted = CryptoJS.AES.encrypt(privateKey, DECRYPTION_KEY);
    return encrypted.toString();
  } catch (error) {
    console.error("Encryption error:", error);
    throw new Error("Failed to encrypt private key");
  }
}

// API Call Helper
async function apiCall(endpoint, method, data = null) {
  try {
    const options = {
      method: method,
      headers: {
        "Content-Type": "application/json",
      },
    };

    if (data) {
      options.body = JSON.stringify(data);
    }

    const response = await fetch(`${API_BASE_URL}${endpoint}`, options);
    const result = await response.json();

    if (!response.ok) {
      throw new Error(result.error || `HTTP error! status: ${response.status}`);
    }

    return result;
  } catch (error) {
    console.error("API call error:", error);
    throw error;
  }
}

// Show result
function showResult(resultId, success, data) {
  const resultDiv = document.getElementById(resultId);
  resultDiv.style.display = "block";
  resultDiv.className = `result ${success ? "success" : "error"}`;

  if (success) {
    let html = "<h3>✅ Success!</h3>";
    for (const [key, value] of Object.entries(data)) {
      if (value) {
        html += `<div class="result-item">
                    <span class="result-label">${formatLabel(
                      key
                    )}:</span> ${value}
                </div>`;
      }
    }
    resultDiv.innerHTML = html;
  } else {
    resultDiv.innerHTML = `<h3>❌ Error</h3><div class="result-item">${data}</div>`;
  }

  // Scroll to result
  resultDiv.scrollIntoView({ behavior: "smooth", block: "nearest" });
}

// Format label
function formatLabel(key) {
  return key.replace(/_/g, " ").replace(/\b\w/g, (l) => l.toUpperCase());
}

// Toggle loading state
function setLoading(button, loading) {
  const btnText = button.querySelector(".btn-text");
  const spinner = button.querySelector(".spinner");

  if (loading) {
    button.disabled = true;
    btnText.style.display = "none";
    spinner.style.display = "inline-block";
  } else {
    button.disabled = false;
    btnText.style.display = "inline";
    spinner.style.display = "none";
  }
}

// Tab switching for operations
document.querySelectorAll(".tab-button").forEach((button) => {
  button.addEventListener("click", () => {
    const tabName = button.getAttribute("data-tab");

    // Update button states
    document
      .querySelectorAll(".tab-button")
      .forEach((btn) => btn.classList.remove("active"));
    button.classList.add("active");

    // Update tab content
    document
      .querySelectorAll(".tab-content")
      .forEach((content) => content.classList.remove("active"));
    document.getElementById(`${tabName}-tab`).classList.add("active");
  });
});

// Tab switching for swap section
document.querySelectorAll(".swap-tab-button").forEach((button) => {
  button.addEventListener("click", () => {
    const tabName = button.getAttribute("data-swap-tab");

    // Update button states
    document
      .querySelectorAll(".swap-tab-button")
      .forEach((btn) => btn.classList.remove("active"));
    button.classList.add("active");

    // Update tab content
    document
      .querySelectorAll(".swap-tab-content")
      .forEach((content) => content.classList.remove("active"));
    document.getElementById(`${tabName}-tab`).classList.add("active");
  });
});

// Load contract information
async function loadContractInfo(contractKey) {
  const contract = CONTRACTS[contractKey];
  const boxElement = document.getElementById(`${contractKey}-contract`);
  const loadingElement = boxElement.querySelector(".contract-loading");
  const dataElement = boxElement.querySelector(".contract-data");
  const errorElement = boxElement.querySelector(".contract-error");

  try {
    // Show loading
    loadingElement.style.display = "flex";
    dataElement.style.display = "none";
    errorElement.style.display = "none";

    // Call API
    const result = await apiCall("/token/contract-address-info", "POST", {
      contract_address: contract.address,
    });

    // Hide loading
    loadingElement.style.display = "none";

    // Display data
    dataElement.style.display = "block";
    dataElement.innerHTML = generateContractInfoHTML(result, contract);
  } catch (error) {
    // Hide loading
    loadingElement.style.display = "none";

    // Show error
    errorElement.style.display = "block";
    errorElement.innerHTML = `
      <h3>❌ Error</h3>
      <p>${error.message}</p>
    `;
  }
}

// Generate HTML for contract information
function generateContractInfoHTML(data, contract) {
  let html = "";

  // Create rows for each data field
  const fields = [
    { key: "address", label: "📍 Address", value: contract.address },
    { key: "name", label: "📛 Name", value: data.name },
    { key: "symbol", label: "🔤 Symbol", value: data.symbol },
    { key: "decimals", label: "🔢 Decimals", value: data.decimals },
    { key: "total_supply", label: "💰 Total Supply", value: data.total_supply },
    {
      key: "owner_address",
      label: "👤 Owner",
      value: data.owner_address,
    },
    {
      key: "owner_balance",
      label: "💵 Owner Balance",
      value: data.owner_balance,
    },
  ];

  fields.forEach((field) => {
    if (
      field.value !== undefined &&
      field.value !== null &&
      field.value !== ""
    ) {
      html += `
        <div class="contract-info-row">
          <div class="contract-info-label">${field.label}</div>
          <div class="contract-info-value">${field.value}</div>
        </div>
      `;
    }
  });

  return html;
}

// Load all contract information on page load
window.addEventListener("DOMContentLoaded", () => {
  console.log("Loading contract information...");

  // Load info for all contracts
  Object.keys(CONTRACTS).forEach((contractKey) => {
    loadContractInfo(contractKey);
  });
});

// Mint Form
document.getElementById("mint-form").addEventListener("submit", async (e) => {
  e.preventDefault();
  const button = e.target.querySelector('button[type="submit"]');
  setLoading(button, true);

  try {
    const contractAddress = document
      .getElementById("mint-contract")
      .value.trim();
    const to = document.getElementById("mint-to").value.trim();
    const amount = document.getElementById("mint-amount").value.trim();
    const privateKey = document.getElementById("mint-private-key").value.trim();

    // Encrypt private key
    const encryptedPrivateKey = encryptPrivateKey(privateKey);

    // Call API
    const result = await apiCall("/token/mint", "POST", {
      contract_address: contractAddress,
      to: to,
      amount: amount,
      encrypted_private_key: encryptedPrivateKey,
    });

    showResult("mint-result", true, result);

    // Clear private key field for security
    document.getElementById("mint-private-key").value = "";
  } catch (error) {
    showResult("mint-result", false, error.message);
  } finally {
    setLoading(button, false);
  }
});

// Burn Form
document.getElementById("burn-form").addEventListener("submit", async (e) => {
  e.preventDefault();
  const button = e.target.querySelector('button[type="submit"]');
  setLoading(button, true);

  try {
    const contractAddress = document
      .getElementById("burn-contract")
      .value.trim();
    const amount = document.getElementById("burn-amount").value.trim();
    const privateKey = document.getElementById("burn-private-key").value.trim();

    // Encrypt private key
    const encryptedPrivateKey = encryptPrivateKey(privateKey);

    // Call API
    const result = await apiCall("/token/burn", "POST", {
      contract_address: contractAddress,
      amount: amount,
      encrypted_private_key: encryptedPrivateKey,
    });

    showResult("burn-result", true, result);

    // Clear private key field for security
    document.getElementById("burn-private-key").value = "";
  } catch (error) {
    showResult("burn-result", false, error.message);
  } finally {
    setLoading(button, false);
  }
});

// Transfer Form
document
  .getElementById("transfer-form")
  .addEventListener("submit", async (e) => {
    e.preventDefault();
    const button = e.target.querySelector('button[type="submit"]');
    setLoading(button, true);

    try {
      const contractAddress = document
        .getElementById("transfer-contract")
        .value.trim();
      const to = document.getElementById("transfer-to").value.trim();
      const amount = document.getElementById("transfer-amount").value.trim();
      const privateKey = document
        .getElementById("transfer-private-key")
        .value.trim();

      // Encrypt private key
      const encryptedPrivateKey = encryptPrivateKey(privateKey);

      // Call API
      const result = await apiCall("/token/transfer", "POST", {
        contract_address: contractAddress,
        to: to,
        amount: amount,
        encrypted_private_key: encryptedPrivateKey,
      });

      showResult("transfer-result", true, result);

      // Clear private key field for security
      document.getElementById("transfer-private-key").value = "";
    } catch (error) {
      showResult("transfer-result", false, error.message);
    } finally {
      setLoading(button, false);
    }
  });

// Balance Form
document
  .getElementById("balance-form")
  .addEventListener("submit", async (e) => {
    e.preventDefault();
    const button = e.target.querySelector('button[type="submit"]');
    setLoading(button, true);

    try {
      const contractAddress = document
        .getElementById("balance-contract")
        .value.trim();
      const address = document.getElementById("balance-address").value.trim();

      // Call API (GET request with query parameters)
      const result = await apiCall("/token/balance", "POST", {
        contract_address: contractAddress,
        address: address,
      });

      showResult("balance-result", true, result);
    } catch (error) {
      showResult("balance-result", false, error.message);
    } finally {
      setLoading(button, false);
    }
  });

// Pre-fill demo data (for testing)
window.fillDemoData = function () {
  // Example VNDX contract address
  document.getElementById("mint-contract").value =
    "0x329aaF4e8d9883c6F8610D48172DE9c6C0917ecD";
  document.getElementById("mint-to").value =
    "0x713641d21610ef2bf6ac7fcf69ae561f6d2ea842";
  document.getElementById("mint-amount").value = "2";
  document.getElementById("mint-private-key").value =
    "e9b1d63e8acd7fe676acb43afb390d4b0202dab61abec9cf2a561e4becb147de";
};

// Swap Execute Form
document
  .getElementById("swap-execute-form")
  .addEventListener("submit", async (e) => {
    e.preventDefault();
    const button = e.target.querySelector('button[type="submit"]');
    setLoading(button, true);

    try {
      const contractAddress = document
        .getElementById("swap-contract")
        .value.trim();
      const amountIn = document.getElementById("swap-amount").value.trim();
      const direction = document.getElementById("swap-direction").value;
      const privateKey = document
        .getElementById("swap-private-key")
        .value.trim();

      // Encrypt private key
      const encryptedPrivateKey = encryptPrivateKey(privateKey);

      // Call API
      const result = await apiCall("/swap", "POST", {
        contract_address: contractAddress,
        amount_in: amountIn,
        direction: direction,
        encrypted_private_key: encryptedPrivateKey,
      });

      showResult("swap-execute-result", true, result);

      // Clear private key field for security
      document.getElementById("swap-private-key").value = "";
    } catch (error) {
      showResult("swap-execute-result", false, error.message);
    } finally {
      setLoading(button, false);
    }
  });

// Swap Quote Form
document
  .getElementById("swap-quote-form")
  .addEventListener("submit", async (e) => {
    e.preventDefault();
    const button = e.target.querySelector('button[type="submit"]');
    setLoading(button, true);

    try {
      const contractAddress = document
        .getElementById("quote-contract")
        .value.trim();
      const amountIn = document.getElementById("quote-amount").value.trim();
      const direction = document.getElementById("quote-direction").value;

      // Call API
      const result = await apiCall("/swap/quote", "POST", {
        contract_address: contractAddress,
        amount_in: amountIn,
        direction: direction,
      });

      showResult("swap-quote-result", true, result);
    } catch (error) {
      showResult("swap-quote-result", false, error.message);
    } finally {
      setLoading(button, false);
    }
  });

// Swap Info Form
document
  .getElementById("swap-info-form")
  .addEventListener("submit", async (e) => {
    e.preventDefault();
    const button = e.target.querySelector('button[type="submit"]');
    setLoading(button, true);

    try {
      const contractAddress = document
        .getElementById("info-contract")
        .value.trim();

      // Call API
      const result = await apiCall("/swap/info", "POST", {
        contract_address: contractAddress,
      });

      showResult("swap-info-result", true, result);
    } catch (error) {
      showResult("swap-info-result", false, error.message);
    } finally {
      setLoading(button, false);
    }
  });

// Add demo button in console
console.log(
  "%c🚀 Token Management Dashboard",
  "font-size: 20px; font-weight: bold; color: #667eea;"
);
console.log(
  "%cTip: Run fillDemoData() to pre-fill the mint form with test data",
  "color: #666;"
);
console.log("%cAPI Base URL:", API_BASE_URL, "color: #43e97b;");
console.log("%cDecryption Key:", DECRYPTION_KEY, "color: #f5576c;");
