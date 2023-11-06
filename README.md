# notsybil

`notsybil` is a command-line interface (CLI) application designed to help you manage cryptocurrency trading on a supported exchange. This bot provides two main commands: "setup" and "withdraw" to configure your API credentials and initiate cryptocurrency withdrawals, respectively.

## Supported Exchange

`notsybil` is currently implemented to work with the OKX exchange. Please note that it may not be compatible with other exchanges at this time.

## API Key Setup (For OKX)

To use the bot with the OKX exchange, you'll need to create API keys with withdrawal permissions. Here are the steps to generate API keys for OKX:

1. Visit [OKX API Management](https://www.okx.com/account/my-api).

2. Log in to your OKX account.

3. Create a new API key.

4. During the API key creation, ensure you grant it the necessary permissions, specifically the "withdrawal" permission.

5. Once the API key is generated, take note of the API Key, Secret Key, and Passphrase.

## Setup
The "setup" command is used to configure your API credentials. Before you can use the bot to withdraw funds, you need to provide your API key, secret key, and passphrase. Here's how to use the "setup" command:

```bash
$ notsybil setup
```

The bot will prompt you to enter your API key, secret key, and passphrase. Once you've provided these credentials, the "setup" command will create a `config.json` file and store these values in it. This config.json file will be used by the bot for subsequent operations, such as withdrawals.
Please ensure that the config.json file is kept secure and not shared with anyone, as it contains sensitive information required for API access.

## Withdraw

The "withdraw" command allows you to initiate cryptocurrency withdrawals from your exchange account. Before using this command, you need to create a `withdraw.csv` file and provide the withdrawal details in the following format:

```
amount,currency,chain,address
```

For example:

```
0.0220,eth,ethereum,0x0730adc99f699e6ddd3d4ffb500b2b27c8113f63ka84b933607be801f7a20ef3
0.0076,eth,starknet,0x0730adc99f699e6ddd3d4ffb500b2b27c8113f63ka84b933607be801f7a20ef3
1.0081,op,optimism,0x0730adc99f699e6ddd3d4ffb500b2b27c8113f63ka84b933607be801f7a20ef3
0.9043,arb,arbitrum,0x0730adc99f699e6ddd3d4ffb500b2b27c8113f63ka84b933607be801f7a20ef3
```

Each line in the `withdraw.csv` file represents a single withdrawal operation with the specified amount, currency, blockchain or chain, and withdrawal address.
Once you've prepared the `withdraw.csv` file with your withdrawal details, you can use the "withdraw" command as follows:

```bash
$ notsybil withdraw
```

When you run the "withdraw" command, the bot will display your available funds and any fees associated with withdrawals. You can confirm or reject each withdrawal operation by typing "y" or "n" when prompted. If you confirm a withdrawal, the bot will proceed with the withdrawal process and display a success message along with the withdrawal ID.

Please make sure to use the "withdraw" command responsibly and ensure that you are following the exchange's rules and policies regarding withdrawals.

## Supported Assets and Networks

`notsybil` currently supports withdrawals for the following assets across multiple networks:

- eth (etereum, arbitrum-one, zksync-era, starknet, optimism, linea)
- usdt (etereum, tron, arbitrum-one, polygon, optimism)
- usdc (etereum, tron, arbitrum-one, polygon, optimism)
- op (optimism)
- arb (arbitrum)
- apt (aptos)

## License
This project is open-source and available under the MIT License.

For more details on how to use this bot, refer to the official documentation or consult the author for support and assistance.
