/**
 * Formats money value
 * @param {string|number} value
 * @return {string}
 */
export const formatMoney = value => {
  if (value === 0 || value === "0") return value;
  if (!value) return value;
  value = Math.round(value * 100) / 100;
  let money = value.toString();
  money = money.replace(/\B(?=(\d{3})+(?!\d))/g, ",");
  return money;
};

/**
 * Formats money value to Naira
 * @param {string|number} value
 * @return {string}
 */
export const formatDollar = value => `$${formatMoney(value)}`;

/**
 * Handles response data by showing a toast on the UI
 * @param {Promise} resp the API response
 */
export const handleResponse = resp => {
  if (resp.status === 200 || resp.status === 201) {
    toast.success(resp.message);
  } else {
    toast.info(resp.message);
  }
};

/**
 * Handles validation error
 * @param {object} error the Joi error object
 * @param {object} obj the error object to populate after checking for the errors
 */
export const handleValidationError = (error, obj) => {
  obj[error.details[0].context.key] = error.details[0].message;
};
