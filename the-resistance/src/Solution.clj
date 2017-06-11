(ns Solution
  (:gen-class))

(def morse-table {\A ".-" \B "-..." \C "-.-." \D "-.."
                  \E ".-" \F "..-." \G "--." \H  "...."
                  \I "..-." \J ".---" \K "-.-." \L ".-.."
                  \M "--." \N "-..." \O "---" \P ".--."
                  \Q "--.-" \R ".-.." \S "..." \T "-"
                  \U "..-." \V "...-" \W ".---" \X "-..-"
                  \Y "-.--" \Z "--.."})

(defn morse
  ([text]
   (morse text ""))
  ([text sep]
   (clojure.string/join sep (map #(get morse-table %) text))))

(defn log
  [msg]
  (binding [*out* *err*]
    (println msg)))

(defn possible-starting-words
  "Given a morse sequence string, and a morse encoded collection of words,
  returns all words that may start the given sequence."
  [morse-sequence morse-words]
  (reduce (fn [starting-words curr]
            (if (clojure.string/starts-with? morse-sequence curr)
              (conj starting-words curr)
              starting-words))
          []
          morse-words))

(defn remove-starting-word
  [morse-sequence word]
  (clojure.string/replace-first morse-sequence
                               (re-pattern word)
                               ""))

(defn possible-word-sequences
  [morse-sequence morse-words]
  (let [starting-words (possible-starting-words morse-sequence morse-words)]
    (->> starting-words
         (map (fn [word]
                 (let [new-morse-sequence (remove-starting-word morse-sequence word)
                       child-messages (possible-word-sequences new-morse-sequence morse-words)]
                   (if (zero? child-messages)
                     1
                     child-messages)
                     )))
        (reduce +))))

(defn number-of-possible-messages
  [morse-sequence dictionary-words]
  (let [morse-words (map morse dictionary-words)]
    (cond
      (= (first morse-words) morse-sequence) 1
      (= (clojure.string/join "" morse-words) morse-sequence) 1
      :else 0)
    ))

(defn load-dictionary-from-stdin
  [dictionary-size]
  (loop [i dictionary-size
         words []]
    (if (= i 0)
      words
      (let [W (read)]
        (recur (dec i) (conj words W))))))

(defn -main [& args]
  (let [L (read) N (read)]
    (println (number-of-possible-messages L (load-dictionary-from-stdin N)))))
