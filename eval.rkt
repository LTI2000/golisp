#lang r5rs

(define-syntax defun
  (syntax-rules ()
    ((defun name params body) (define (name . params) body))))

(define (atom x) (not (pair? x)))

(define eq eq?)

(defun null. (x)
  (eq x '()))

(defun and. (x y)
  (cond (x (cond (y 't) ('t #false)))
        ('t #false)))

(defun not. (x)
  (cond ((null. x) 't)
        (x #false)
        ('t 't)))

(defun append. (x y)
  (cond ((null. x) y)
        ('t (cons (car x) (append. (cdr x) y)))))

(defun pair. (x y)
  (cond ((and. (null. x) (null. y)) '())
        ((and. (not. (atom x)) (not. (atom y)))
         (cons (list (car x) (car y))
               (pair. (cdr x) (cdr y))))))

(defun assoc. (x y)
  (cond ((eq (caar y) x) (cadar y))
        ('t (assoc. x (cdr y)))))

(defun eval. (e a)
  (cond
    ((atom e) (assoc. e a))
    ((atom (car e))
     (cond
       ((eq (car e) 'quote) (cadr e))
       ((eq (car e) 'atom)  (atom (eval. (cadr e) a)))
       ((eq (car e) 'eq)    (eq   (eval. (cadr e) a)
                                  (eval. (caddr e) a)))
       ((eq (car e) 'car)   (car  (eval. (cadr e) a)))
       ((eq (car e) 'cdr)   (cdr  (eval. (cadr e) a)))
       ((eq (car e) 'cons)  (cons (eval. (cadr e) a)
                                  (eval. (caddr e) a)))
       ((eq (car e) 'cond)  (evcon. (cdr e) a))
       ('t (eval. (cons (assoc. (car e) a)
                        (cdr e))
                  a))))
    ((eq (caar e) 'label)
     (eval. (cons (caddar e) (cdr e))
            (cons (list (cadar e) (car e)) a)))
    ((eq (caar e) 'lambda)
     (eval. (caddar e)
            (append. (pair. (cadar e) (evlis. (cdr e) a))
                     a)))))

(defun evcon. (c a)
  (cond ((eval. (caar c) a)
         (eval. (cadar c) a))
        ('t (evcon. (cdr c) a))))

(defun evlis. (m a)
  (cond ((null. m) '())
        ('t (cons (eval.  (car m) a)
                  (evlis. (cdr m) a)))))


(display (eval. '(subst 'm 'b '(a b (a b c) d))
                '((subst (lambda (x y z)
                           (cond ((atom z)
                                  (cond ((eq z y) x)
                                        ('t z)))
                                 ('t (cons (subst x y (car z))
                                           (subst x y (cdr z))))))))))
(display "\n")

(display (eval. '((label subst 
                         (lambda (x y z)
                           (cond ((atom z)
                                  (cond ((eq z y) x)
                                        ('t z)))
                                 ('t (cons (subst x y (car z))
                                           (subst x y (cdr z))))))) 'm 'b '(a b (a b c) d)) '()))
(display "\n")
