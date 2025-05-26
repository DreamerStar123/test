package com.example;

import com.google.crypto.tink.apps.paymentmethodtoken.GooglePaymentsPublicKeysManager;
import com.google.crypto.tink.apps.paymentmethodtoken.PaymentMethodTokenRecipient;
import io.github.cdimascio.dotenv.Dotenv;

import java.security.KeyFactory;
import java.security.spec.PKCS8EncodedKeySpec;
import java.security.PrivateKey;
import java.util.Base64;

/**
 * mvn compile
 * mvn exec:java
 * mvn exec:java -Dexec.mainClass="com.example.App"
 */

public class App {

    // Converts a Base64 PKCS#8 private key string to a PrivateKey object
    private static PrivateKey loadPrivateKey(String base64Key) throws Exception {
        byte[] keyBytes = Base64.getDecoder().decode(base64Key);
        PKCS8EncodedKeySpec keySpec = new PKCS8EncodedKeySpec(keyBytes);
        KeyFactory kf = KeyFactory.getInstance("EC");  // Elliptic Curve
        return kf.generatePrivate(keySpec);
    }

    public static void main(String[] args) throws Exception {
        // Optional but recommended to prefetch public signing keys
        GooglePaymentsPublicKeysManager.INSTANCE_PRODUCTION.refreshInBackground();

        // Replace with your actual gateway ID registered with Google
        String gatewayId = "gateway:stripe";

        // Replace with your private key (base64 PKCS#8 format)
        String base64PrivateKey = "MIGHAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBG0wawIBAQQg0UpNwizASg7jrxOlzeIs/6stm2/bUAwodTw/XVE1UXqhRANCAAQMhlVfSYafgp3KSZJvWeEd188YWZYTF9N3IVrL8Oz93JSfFBaMDaJGD3EIbo13gLTEU9lvmTGiqSqMRfxbU9eR"; // truncated for brevity
        PrivateKey privateKey = loadPrivateKey(base64PrivateKey);

        // Load the encrypted token from a local .env file
        Dotenv dotenv = Dotenv.load();
        String encryptedToken = dotenv.get("ENCRYPTED_TOKEN2");
        if (encryptedToken == null || encryptedToken.isEmpty()) {
            throw new IllegalArgumentException("ENCRYPTED_TOKEN is not set or empty in the .env file.");
        }

        String decryptedJson = new PaymentMethodTokenRecipient.Builder()
            .fetchSenderVerifyingKeysWith(GooglePaymentsPublicKeysManager.INSTANCE_PRODUCTION)
            .recipientId(gatewayId)
            .protocolVersion("ECv2")
            .addRecipientPrivateKey(base64PrivateKey)
            .build()
            .unseal(encryptedToken);

        System.out.println("Decrypted token: " + decryptedJson);
    }
}
