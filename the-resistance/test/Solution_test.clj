(ns Solution-test
  (:require [clojure.test :refer :all]
            [Solution :as solution]))

(def morse solution/morse)

(deftest example-test
  (let [morse-sequence "......-...-..---.-----.-..-..-.."
        dictionary ["HELL"
                    "HELLO"
                    "OWORLD"
                    "WORLD"
                    "TEST"]]
    (is (= 2 (solution/possible-word-sequences morse-sequence (map morse dictionary) false)))))

(deftest test-01
  (let [morse-sequence "-.-"
        dictionary ["A"
                    "B"
                    "C"
                    "HELLO"
                    "K"
                    "WORLD"]]
    (is (= 1 (solution/possible-word-sequences morse-sequence (map morse dictionary))))))

(deftest test-02
  (let [morse-sequence "--.-------.."
        dictionary ["GOD"
                    "GOOD"
                    "MORNING"
                    "G"
                    "HELLO"]]
    (is (= 1 (solution/possible-word-sequences morse-sequence (map morse dictionary) false)))))

(deftest test-03
  (let [morse-sequence "......-...-..---.-----.-..-..-.."
        dictionary ["HELL"
                    "HELLO"
                    "OWORLD"
                    "WORLD"
                    "TEST"]]
    (is (= 2 (solution/possible-word-sequences morse-sequence (map morse dictionary) false)))))

(deftest simplified-example-HELLO-test
  (let [morse-sequence (solution/morse "HELLO")
        dictionary ["HELLO"]]
    (is (= 1 (solution/number-of-possible-messages morse-sequence dictionary)))))

(deftest simplified-example-HELLO-WORLD-test
  (let [morse-sequence (solution/morse "HELLOWORLD")
        dictionary ["HELLO" "WORLD"]]
    (is (= 1 (solution/number-of-possible-messages morse-sequence dictionary)))))

#_(deftest simplified-example-HELLO-WORLD-test
  (let [morse-sequence (solution/morse "HELLOWORLD")
        dictionary ["HELLO" "WORLD" "HELL" "OWORLD"]]
    (is (= 2 (solution/number-of-possible-messages morse-sequence dictionary)))))

(deftest spaceless-message-test
  (is (= (morse "HELLO") "......-...-..---"))
  (is (= (morse "HELLO" " ") ".... . .-.. .-.. ---")))

(deftest possible-starting-words-test
  (is (= [(morse "HELL")]
         (solution/possible-starting-words
           (morse "HELLO")
           [(morse "HELL")])))
  (is (= [(morse "HELLO") (morse "HELL")]
         (solution/possible-starting-words
           (morse "HELLOBABE")
           [(morse "HELLO") (morse "HELL")]))))

(defn test-possible-n [expected-n text dictionary]
  (is (= expected-n
         (solution/possible-word-sequences
           (morse text)
           (map morse dictionary)))))

(deftest possible-words-test
  (testing "simple sequence"
    (test-possible-n 1 "HELLO" ["HELLO"])
    (test-possible-n 1 "HELLOWORLD" ["HELLO" "WORLD" "TEST"])
    (test-possible-n 1 "HELLOWORLDTEST" ["HELLO" "WORLD" "TEST"]))
  (testing "simple branch"
    (test-possible-n 2 "HELLOWORLD" ["HELLO" "WORLD" "HELL" "OWORLD"])))
