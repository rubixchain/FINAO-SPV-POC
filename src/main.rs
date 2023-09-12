use umbral_pre::*;

// As in any public-key cryptosystem, users need a pair of public and private keys.
// Additionally, users that delegate access to their data (like Alice, in this example)
// need a signing keypair.

// Key Generation (on Alice's side)


struct KeyPairBytes
{
    pub_key: [u8; 33],
    pvt_key: [u8; 32],
    sign_key: [u8; 32],
    ver_key: [u8; 33],
}

struct KeyFragBytes
{
    keyfrag0: [u8; 260],
    keyfrag1: [u8; 260],
    capsule:  [u8; 98],
    ct: Box<[u8]>,
}

fn main() {

    let alice = create_key_pair();
    let bob = create_key_pair();

    let _bytes: &[u8] = unsafe { any_as_u8_slice(&alice) };

    let enc_data = encrypt_data(b"testing value ", alice.pub_key, alice.pvt_key, bob.pub_key, alice.sign_key);
    decrypt_data(enc_data.keyfrag0, enc_data.keyfrag1, enc_data.capsule, alice.ver_key, alice.pub_key, bob.pub_key, bob.pvt_key , enc_data.ct)

}

fn create_key_pair() -> KeyPairBytes
{
let secret_key = SecretKey::random();
let public_key = secret_key.public_key();
let signing_key = SecretKey::random();
let _signer = Signer::new(&signing_key);
let verifying_pk = signing_key.public_key();

let pvt_key_array = secret_key.to_array();
let pvt_key_bytes: [u8; 32] = pvt_key_array.as_slice().try_into().expect("Wrong length");

let pub_key_array = public_key.to_array();
let pub_key_bytes: [u8; 33] = pub_key_array.as_slice().try_into().expect("Wrong length");

let signing_key_array = signing_key.to_array();
let signing_key_bytes: [u8; 32] = signing_key_array.as_slice().try_into().expect("Wrong length");

let ver_key_array = verifying_pk.to_array();
let ver_key_bytes: [u8; 33] = ver_key_array.as_slice().try_into().expect("Wrong length");

let key_pair_generated = KeyPairBytes {
    pub_key: pub_key_bytes,
    pvt_key: pvt_key_bytes,
    sign_key: signing_key_bytes,
    ver_key: ver_key_bytes,
};

return key_pair_generated;

}

fn encrypt_data( plaintext: &[u8;14 ],pub_key_bytes: [u8; 33], pvt_key_bytes:[u8; 32], delegator_pub_key: [u8; 33], sign_key_bytes:[u8; 32] ) -> KeyFragBytes
{
    // let plaintext = b"peace at dawn";
    let pub_key_from_bytes = PublicKey::from_bytes(pub_key_bytes).unwrap();
    let pvt_key_from_bytes = SecretKey::from_bytes(pvt_key_bytes).unwrap();

    let sign_key_from_bytes = SecretKey::from_bytes(sign_key_bytes).unwrap();

    let signer = Signer::new(&sign_key_from_bytes);

    let delegator_pub_key_from_bytes = PublicKey::from_bytes(delegator_pub_key).unwrap();

let (capsule, ciphertext) = encrypt(&pub_key_from_bytes, plaintext).unwrap();
let plaintext_owner = decrypt_original(&pvt_key_from_bytes, &capsule, &ciphertext).unwrap();
assert_eq!(&plaintext_owner as &[u8], plaintext);

let n = 2; // how many fragments to create
let m = 2; // how many should be enough to decrypt
let verified_kfrags = generate_kfrags(&pvt_key_from_bytes, &delegator_pub_key_from_bytes, &signer, m, n, true, true);

let kfrag0 = KeyFrag::from_array(&verified_kfrags[0].to_array()).unwrap();
let kfrag1 = KeyFrag::from_array(&verified_kfrags[1].to_array()).unwrap();


let kfrag0_array = kfrag0.to_array();
let kfrag1_array = kfrag1.to_array();
let kfrag0_bytes: [u8; 260] = kfrag0_array.as_slice().try_into().expect("Wrong length");
let kfrag1_bytes: [u8; 260] = kfrag1_array.as_slice().try_into().expect("Wrong length");


let capsule_array = capsule.to_array();
let capsule_bytes: [u8; 98] = capsule_array.as_slice().try_into().expect("Wrong length");

let key_frag = KeyFragBytes {
    keyfrag0:kfrag0_bytes,
    keyfrag1:kfrag1_bytes,
    capsule: capsule_bytes,
    ct: ciphertext,
};

return key_frag;


}

fn decrypt_data(kfrag0_bytes:[u8; 260], kfrag1_bytes:[u8; 260], capsule_bytes: [u8; 98], ver_key_bytes: [u8; 33], owner_pub_key_bytes: [u8; 33], delegator_pub_key_bytes: [u8; 33] , delegator_pvt_key_bytes: [u8; 32], ciphertext: Box<[u8]>)
{
    let kfrag0 = KeyFrag::from_bytes(kfrag0_bytes).unwrap();
    let kfrag1 = KeyFrag::from_bytes(kfrag1_bytes).unwrap();
    let capsule = Capsule::from_bytes(capsule_bytes).unwrap();
    let verifying_pk = PublicKey::from_bytes(ver_key_bytes).unwrap();
    let owner_pk = PublicKey::from_bytes(owner_pub_key_bytes).unwrap();
    let delegator_pk = PublicKey::from_bytes(delegator_pub_key_bytes).unwrap();
    let delegator_pvt_key = SecretKey::from_bytes(delegator_pvt_key_bytes).unwrap();

    // // Ursula 0
let verified_kfrag0 = kfrag0.verify(&verifying_pk, Some(&owner_pk), Some(&delegator_pk)).unwrap();
let verified_cfrag0 = reencrypt(&capsule, &verified_kfrag0);

// // Ursula 1
let verified_kfrag1 = kfrag1.verify(&verifying_pk, Some(&owner_pk), Some(&delegator_pk)).unwrap();
let verified_cfrag1 = reencrypt(&capsule, &verified_kfrag1);

let cfrag0 = CapsuleFrag::from_array(&verified_cfrag0.to_array()).unwrap();
let cfrag1 = CapsuleFrag::from_array(&verified_cfrag1.to_array()).unwrap();

let verified_cfrag0 = cfrag0
    .verify(&capsule, &verifying_pk, &owner_pk, &delegator_pk)
    .unwrap();
let verified_cfrag1 = cfrag1
    .verify(&capsule, &verifying_pk, &owner_pk, &delegator_pk)
    .unwrap();

    let plaintext_delegrator = decrypt_reencrypted(
    &delegator_pvt_key, &owner_pk, &capsule, &[verified_cfrag0, verified_cfrag1], &ciphertext).unwrap();

let plaintext_decrypted = String::from_utf8(plaintext_delegrator.to_vec()).unwrap();

println!("final plaintext {:?}" , plaintext_decrypted);

}
unsafe fn any_as_u8_slice<T: Sized>(p: &T) -> &[u8] {
    ::core::slice::from_raw_parts(
        (p as *const T) as *const u8,
        ::core::mem::size_of::<T>(),
    )
}