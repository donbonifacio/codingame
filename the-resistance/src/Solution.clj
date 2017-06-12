(ns Solution
  (:gen-class))

(def morse-table {\A ".-" \B "-..." \C "-.-." \D "-.."
                  \E "." \F "..-." \G "--." \H  "...."
                  \I ".." \J ".---" \K "-.-" \L ".-.."
                  \M "--" \N "-." \O "---" \P ".--."
                  \Q "--.-" \R ".-." \S "..." \T "-"
                  \U "..-" \V "...-" \W ".--" \X "-..-"
                  \Y "-.--" \Z "--.."})

(defn morse
  ([text]
   (morse text ""))
  ([text sep]
   (clojure.string/join sep (map #(get morse-table %) (str text)))))

(defn log
  [msg]
  (binding [*out* *err*]
    (println msg)))

(defn starts-with? [s needle]
  (.startsWith (.toString s) needle))

(defn possible-starting-words
  "Given a morse sequence string, and a morse encoded collection of words,
  returns all words that may start the given sequence."
  [morse-sequence morse-words]
  (reduce (fn [starting-words curr]
            (if (starts-with? morse-sequence curr)
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
  ([morse-sequence morse-words]
   (possible-word-sequences morse-sequence morse-words false))
  ([morse-sequence morse-words verbose]
   (let [starting-words (possible-starting-words morse-sequence morse-words)]
     (when verbose
       (println "Sequence: " morse-sequence)
       (println "Starting words: " starting-words))
     (->> starting-words
          (map (fn [word]
                  (let [new-morse-sequence (remove-starting-word morse-sequence word)
                        child-messages (possible-word-sequences new-morse-sequence morse-words verbose)]
                    (when verbose
                      (println word " -> " new-morse-sequence " : " child-messages))
                    (cond
                      ;; no childs down the line that match
                      (and (zero? child-messages) (not (empty? new-morse-sequence))) 0
                      ;; has a complete child tree
                      (zero? child-messages) 1
                      ;; has several complete child trees
                      :else child-messages)
                      )))
         (reduce +)))))

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
  (let [L (read)
        N (read)]
    (println (possible-word-sequences
               L
               (map morse (load-dictionary-from-stdin N))))))
