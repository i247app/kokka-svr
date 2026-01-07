// Configuration
const DECRYPTION_KEY = "KOK_EVM_ENCRYPTED_8888";
const API_BASE_URL = "https://m1.i247.com/kokka"; // Update this to your server URL

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

// Tab switching
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
      const result = await apiCall(
        `/token/balance?contract_address=${encodeURIComponent(
          contractAddress
        )}&address=${encodeURIComponent(address)}`,
        "GET"
      );

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
